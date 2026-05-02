"""
HTTP3 Client Wrapper API

A Flask-based REST API for executing HTTP3 client commands with logging,
retry logic, and comprehensive monitoring capabilities.
"""

from flask import Flask, jsonify, request, send_from_directory, send_file
import subprocess
import threading
import uuid
import os
import re
import json
import gzip
import tarfile
import tempfile
import shutil
from datetime import datetime, timedelta
import time
import signal
import select
import queue
import multiprocessing
from pathlib import Path
import io

# Initialize Flask application
app = Flask(__name__)

# Global dictionaries for task management
tasks = {}              # Main task storage
cert_tasks = {}         # Certificate retrieval tasks
processes = {}          # Active subprocesses
task_threads = {}       # Thread tracking for async task execution

# ANSI escape sequence regex for cleaning terminal output
ansi_escape = re.compile(r"\x1B(?:[@-Z\\-_]|\[[0-?]*[ -/]*[@-~])")

# Paths to executable binaries
HTTP3_CLIENT = "/usr/app/bin/http3client"
CERTIFICATE_RETRIEVER = "/usr/app/bin/cert-retriever"

# Configuration constants
COMMAND_TIMEOUT = 3600          # Maximum command execution time (1 hour)
LOG_BASE_DIR = "/var/log/http3_client"  # Centralized log directory
LOG_RETENTION_DAYS = 7          # Days to keep logs before cleanup

# Ensure log directory exists
os.makedirs(LOG_BASE_DIR, exist_ok=True)

# Global process pool for command execution
process_pool = None


def init_process_pool():
    """
    Initialize a multiprocessing pool for handling CPU-intensive operations.
    
    This pool is used to prevent blocking the main thread during command execution
    and provides better resource management for concurrent tasks.
    """
    global process_pool
    if process_pool is None:
        process_pool = multiprocessing.Pool(
            processes=4,           # Number of worker processes
            maxtasksperchild=10    # Recycle workers after 10 tasks
        )


def get_log_dir(task_id=None, attempt_id=None):
    """
    Get the appropriate log directory path based on task and attempt IDs.
    
    Directory structure:
    /var/log/http3_client/
    ├── tasks/
    │   ├── {task_id}/
    │   │   ├── attempts/
    │   │   │   ├── {attempt_id}.log.gz
    │   │   │   └── {attempt_id}.meta.json
    │   │   ├── metadata.json
    │   │   └── summary.json
    └── certificates/
        └── {task_id}/
            └── output.json
    
    Args:
        task_id (str, optional): Unique identifier for a task
        attempt_id (str, optional): Unique identifier for an attempt within a task
    
    Returns:
        str: Path to the appropriate log directory
    """
    if task_id:
        task_dir = os.path.join(LOG_BASE_DIR, "tasks", task_id)
        if attempt_id:
            return os.path.join(task_dir, "attempts")
        return task_dir
    return LOG_BASE_DIR


def save_task_log(attempt_id, command, stdout, stderr, metadata=None):
    """
    Save execution logs for a specific attempt in compressed format.
    
    Logs include command execution details, timestamps, stdout, stderr,
    and optional metadata for debugging and analysis.
    
    Args:
        attempt_id (str): Unique identifier for the attempt
        command (str): The command that was executed
        stdout (str): Standard output from the command
        stderr (str): Standard error from the command
        metadata (dict, optional): Additional execution metadata
    
    Returns:
        str: Path to the saved log file
    """
    try:
        # Extract task_id from attempt_id if present
        task_id = attempt_id.split("_attempt_")[0] if "_attempt_" in attempt_id else attempt_id
        attempt_dir = os.path.join(get_log_dir(task_id=task_id), "attempts")
        os.makedirs(attempt_dir, exist_ok=True)
        
        log_file = os.path.join(attempt_dir, f"{attempt_id}.log.gz")
        
        # Write compressed log with structured sections
        with gzip.open(log_file, "wt", encoding="utf-8") as fd:
            fd.write(f"=== LOG FOR ATTEMPT: {attempt_id} ===\n")
            fd.write(f"Timestamp: {datetime.now().isoformat()}\n")
            fd.write(f"Command: {command}\n")
            
            if metadata:
                fd.write(f"\n=== METADATA ===\n")
                for key, value in metadata.items():
                    fd.write(f"{key}: {value}\n")
            
            fd.write(f"\n=== STDOUT ===\n")
            fd.write(stdout)
            
            fd.write(f"\n=== STDERR ===\n")
            fd.write(stderr)
            
            fd.write(f"\n=== END OF LOG ===\n")
        
        # Save metadata as separate JSON file for easy parsing
        if metadata:
            meta_file = os.path.join(attempt_dir, f"{attempt_id}.meta.json")
            with open(meta_file, "w") as f:
                json.dump(metadata, f, indent=2, default=str)
        
        return log_file
        
    except Exception as e:
        # Fallback to temporary directory if main log directory fails
        print(f"Error saving log for attempt {attempt_id}: {e}")
        fallback_file = f"/tmp/output_{attempt_id}.txt.gz"
        with gzip.open(fallback_file, "wt", encoding="utf-8") as fd:
            fd.write(f"Command: {command}\n")
            fd.write(stdout)
            fd.write("\n---------------\n")
            fd.write(stderr)
        return fallback_file


def save_task_metadata(task_id, task_data):
    """
    Save comprehensive metadata for a complete task execution.
    
    This includes task configuration, execution results, and statistical
    summaries for later analysis and reporting.
    
    Args:
        task_id (str): Unique identifier for the task
        task_data (dict): Complete task execution data
    
    Returns:
        str or None: Path to the metadata file, or None on failure
    """
    try:
        task_dir = get_log_dir(task_id=task_id)
        os.makedirs(task_dir, exist_ok=True)
        
        meta_file = os.path.join(task_dir, "metadata.json")
        
        # Create serializable metadata (exclude non-JSON-serializable objects)
        serializable_data = {
            "task_id": task_data.get("task_id"),
            "status": task_data.get("status"),
            "arguments": task_data.get("arguments", []),
            "created_at": task_data.get("created_at"),
            "start_time": task_data.get("start_time"),
            "end_time": task_data.get("end_time"),
            "total_attempts": task_data.get("total_attempts", 0),
            "successful_attempts": task_data.get("successful_attempts", 0),
            "failed_attempts": task_data.get("failed_attempts", 0),
            "total_execution_time": task_data.get("total_execution_time", 0),
            "has_duration": task_data.get("has_duration", False),
            "duration_seconds": task_data.get("duration_seconds"),
            "requested_duration": task_data.get("requested_duration"),
            "achieved_duration_ratio": task_data.get("achieved_duration_ratio"),
            "auto_retry_enabled": task_data.get("auto_retry_enabled", False),
        }
        
        with open(meta_file, "w") as f:
            json.dump(serializable_data, f, indent=2, default=str)
        
        # Save separate summary of attempts for quick analysis
        if task_data.get("attempts"):
            summary_file = os.path.join(task_dir, "summary.json")
            attempts_summary = []
            
            for attempt in task_data["attempts"]:
                attempts_summary.append({
                    "attempt_id": attempt.get("attempt_id"),
                    "status": attempt.get("status"),
                    "execution_time": attempt.get("execution_time"),
                    "returncode": attempt.get("returncode"),
                    "start_time": attempt.get("start_time"),
                    "end_time": attempt.get("end_time"),
                    "has_stats": "stats" in attempt,
                })
            
            with open(summary_file, "w") as f:
                json.dump(attempts_summary, f, indent=2, default=str)
        
        return meta_file
        
    except Exception as e:
        print(f"Error saving metadata for task {task_id}: {e}")
        return None


def cleanup_old_logs():
    """
    Remove log files older than the configured retention period.
    
    This function cleans up the log directory by:
    1. Identifying files older than LOG_RETENTION_DAYS
    2. Removing those files
    3. Removing empty directories
    
    Should be run periodically to prevent disk space exhaustion.
    """
    try:
        cutoff_time = time.time() - (LOG_RETENTION_DAYS * 86400)
        
        for root, dirs, files in os.walk(LOG_BASE_DIR):
            # Remove old files
            for file in files:
                file_path = os.path.join(root, file)
                try:
                    if os.path.getmtime(file_path) < cutoff_time:
                        os.remove(file_path)
                        print(f"Removed old log file: {file_path}")
                except (OSError, PermissionError) as e:
                    print(f"Error removing file {file_path}: {e}")
                    continue
            
            # Remove empty directories (bottom-up)
            for dir_name in dirs:
                dir_path = os.path.join(root, dir_name)
                try:
                    if not os.listdir(dir_path):
                        os.rmdir(dir_path)
                        print(f"Removed empty directory: {dir_path}")
                except (OSError, PermissionError) as e:
                    print(f"Error removing directory {dir_path}: {e}")
                    continue
        
        print(f"Log cleanup completed (retention: {LOG_RETENTION_DAYS} days)")
        
    except Exception as e:
        print(f"Error during log cleanup: {e}")


def create_logs_tarball(task_id=None):
    """
    Create a compressed tarball (.tgz) containing log files.
    
    If task_id is provided, only logs for that specific task are included.
    Otherwise, all logs in the log directory are included.
    
    Args:
        task_id (str, optional): Specific task ID to include, or None for all logs
    
    Returns:
        tuple: (tarball_path, error_message) - path to created tarball or error message
    """
    temp_dir = None
    try:
        # Create temporary working directory
        temp_dir = tempfile.mkdtemp(prefix="http3_logs_")
        
        if task_id:
            # Include only specific task logs
            source_dir = get_log_dir(task_id=task_id)
            if not os.path.exists(source_dir):
                return None, f"No logs found for task {task_id}"
            
            # Copy task logs to temporary directory
            dest_dir = os.path.join(temp_dir, f"task_{task_id}")
            shutil.copytree(source_dir, dest_dir)
            
            tarball_name = f"task_{task_id}_logs_{datetime.now().strftime('%Y%m%d_%H%M%S')}.tgz"
        else:
            # Include all logs
            if not os.path.exists(LOG_BASE_DIR) or not os.listdir(LOG_BASE_DIR):
                return None, "No logs found in log directory"
            
            # Copy entire log structure
            for item in os.listdir(LOG_BASE_DIR):
                src = os.path.join(LOG_BASE_DIR, item)
                dst = os.path.join(temp_dir, item)
                if os.path.isdir(src):
                    shutil.copytree(src, dst)
                else:
                    shutil.copy2(src, dst)
            
            tarball_name = f"all_logs_{datetime.now().strftime('%Y%m%d_%H%M%S')}.tgz"
        
        # Create compressed tarball
        tarball_path = os.path.join(temp_dir, tarball_name)
        
        with tarfile.open(tarball_path, "w:gz") as tar:
            tar.add(temp_dir, arcname=os.path.basename(temp_dir.rstrip('/')))
        
        return tarball_path, None
    
    except Exception as e:
        return None, str(e)
    
    finally:
        # Clean up temporary directory (keeping only the tarball)
        if temp_dir:
            try:
                # Find the created tarball
                tarball_path = None
                for file in os.listdir(temp_dir):
                    if file.endswith('.tgz'):
                        tarball_path = os.path.join(temp_dir, file)
                        break
                
                # Remove everything except the tarball
                for item in os.listdir(temp_dir):
                    item_path = os.path.join(temp_dir, item)
                    if item_path != tarball_path:
                        if os.path.isdir(item_path):
                            shutil.rmtree(item_path)
                        else:
                            os.remove(item_path)
            except Exception as cleanup_error:
                print(f"Error during temporary directory cleanup: {cleanup_error}")


def delete_logs(task_id=None):
    """
    Delete log files from the system.
    
    Args:
        task_id (str, optional): Specific task ID to delete logs for,
                                 or None to delete all logs
    
    Returns:
        tuple: (success, message) - boolean success flag and status message
    """
    try:
        if task_id:
            # Delete logs for specific task
            task_dir = get_log_dir(task_id=task_id)
            if os.path.exists(task_dir):
                shutil.rmtree(task_dir)
                return True, f"Logs for task {task_id} deleted successfully"
            else:
                return False, f"No logs found for task {task_id}"
        else:
            # Delete all logs (but keep directory structure)
            if os.path.exists(LOG_BASE_DIR):
                for item in os.listdir(LOG_BASE_DIR):
                    item_path = os.path.join(LOG_BASE_DIR, item)
                    if os.path.isdir(item_path):
                        shutil.rmtree(item_path)
                    else:
                        os.remove(item_path)
                return True, "All logs deleted successfully"
            else:
                return False, "Log directory not found"
    
    except Exception as e:
        return False, f"Error deleting logs: {str(e)}"


def safe_subprocess_run(cmd, timeout):
    """
    Execute a shell command safely with timeout handling.
    
    This function prevents busy-waiting and properly handles command
    timeouts without consuming excessive CPU resources.
    
    Args:
        cmd (str): Shell command to execute
        timeout (int): Maximum execution time in seconds
    
    Returns:
        dict: Command execution results including return code, stdout, stderr,
              and timeout status
    """
    try:
        result = subprocess.run(
            cmd,
            shell=True,
            timeout=timeout,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True,
            encoding='utf-8',
            errors='ignore'
        )
        return {
            'returncode': result.returncode,
            'stdout': result.stdout,
            'stderr': result.stderr,
            'timeout': False
        }
    except subprocess.TimeoutExpired as e:
        return {
            'returncode': -1,
            'stdout': e.stdout.decode('utf-8', errors='ignore') if e.stdout else '',
            'stderr': e.stderr.decode('utf-8', errors='ignore') if e.stderr else '',
            'timeout': True
        }
    except Exception as e:
        return {
            'returncode': -1,
            'stdout': '',
            'stderr': str(e),
            'timeout': False,
            'error': str(e)
        }


def extract_duration_from_args(arguments):
    """
    Extract duration value from command arguments.
    
    Supports various duration formats:
    - --duration=300
    - --duration=300s (seconds)
    - --duration=5m (minutes)
    - --duration=2h (hours)
    - --duration=1d (days)
    - --duration 300 (space separated)
    
    Args:
        arguments (list): List of command arguments
    
    Returns:
        int or None: Duration in seconds, or None if not found
    """
    args_str = " ".join(arguments) if isinstance(arguments, list) else str(arguments)
    args_str = re.sub(r"\s+", " ", args_str)  # Normalize whitespace
    
    duration_patterns = [
        r"--duration[= ](\d+[smhd]?)(?:\s|$)",  # Matches --duration=value or --duration value
        r"-duration[= ](\d+[smhd]?)(?:\s|$)",   # Matches -duration=value or -duration value
    ]
    
    for pattern in duration_patterns:
        matches = re.findall(pattern, args_str)
        if matches:
            duration_str = matches[0]
            
            # Convert to seconds based on suffix
            if duration_str.endswith("s"):
                return int(duration_str[:-1])
            elif duration_str.endswith("m"):
                return int(duration_str[:-1]) * 60
            elif duration_str.endswith("h"):
                return int(duration_str[:-1]) * 3600
            elif duration_str.endswith("d"):
                return int(duration_str[:-1]) * 86400
            else:
                # No suffix, assume seconds
                try:
                    return int(duration_str)
                except (ValueError, TypeError):
                    return None
    
    return None


def create_attempt_id(task_id, attempt_num):
    """
    Create a unique identifier for an attempt within a task.
    
    Args:
        task_id (str): Parent task identifier
        attempt_num (int): Attempt number (1-based)
    
    Returns:
        str: Formatted attempt identifier
    """
    return f"{task_id}_attempt_{attempt_num}"


def get_counter(regex_pattern, text):
    """
    Extract numerical values from text using regex patterns.
    
    Args:
        regex_pattern (str): Regular expression with capture group
        text (str): Text to search within
    
    Returns:
        int or str: Extracted value, or 0 if not found
    """
    found = re.findall(regex_pattern, text)
    if found:
        try:
            return int(float(found[0]))
        except (ValueError, TypeError):
            return found[0]  # Return as string if not convertible
    else:
        return 0


def get_value(text):
    """
    Parse numerical values with optional unit suffixes (K, M, G, T).
    
    Supports formats like:
    - "1024" (plain number)
    - "1.5K" (kilobytes)
    - "2.0M" (megabytes)
    - "0.5G" (gigabytes)
    
    Args:
        text (str): Text containing a numerical value
    
    Returns:
        float or int or str: Parsed value in base units
    """
    # Check for unit suffixes
    vals = re.findall(r"(\d+\.\d+)([A-Z])", text)
    unit_multipliers = {
        "K": 1024,
        "M": 1024 * 1024,
        "G": 1024 * 1024 * 1024,
        "T": 1024 * 1024 * 1024 * 1024,
    }
    
    if vals:
        number, unit = vals[0]
        return float(number) * unit_multipliers.get(unit, 1)
    else:
        # No unit suffix, try to parse as number
        if "." in text:
            try:
                return float(text)
            except (ValueError, TypeError):
                return text
        else:
            try:
                return int(text)
            except (ValueError, TypeError):
                return text


def process_output(stdout):
    """
    Parse HTTP3/QUIC client output to extract performance metrics.
    
    Extracts key metrics from gladiator client logs including:
    - Throughput statistics
    - Connection counts
    - Packet statistics
    - Error counts
    
    Args:
        stdout (str): Combined stdout and stderr from client execution
    
    Returns:
        dict: Dictionary of extracted metrics with default values of 0
    """
    # Define metric headers for different output formats
    headers = [
        "udp_throughput_rx",
        "udp_throughput_tx",
        "udp_bytes_rx",
        "udp_bytes_tx",
        "inner_connections_created",
        "inner_connections_closed",
        "quic_streams_created",
        "quic_streams_closed",
        "outer_connections_created",
        "outer_connections_closed",
        "quic_connections_created",
        "quic_connections_closed",
    ]
    
    # Regex patterns for specific counters
    counter_regexp = {
        "payload_bytes": r".*Total of.* (\d+\.\d*\w).*bytes of user payload data received in.*\d+\.\d*.*ms",
        "total_time": r".*Total of.*\d+\.\d*.*bytes of user payload data received in.* (\d+\.\d*).*ms",
        "http3_errors": r"There are.* (\d+).*QUIC connections fail to setup with the Target Server.",
        "masque_errors": r"There are.* (\d+).*QUIC connections fail to setup with the MASQUE",
        "tunnel_rx_pkts": r"QUIC-Tunnel packets Rx:.*(\d+).*Tx:.*\d+",
        "tunnel_tx_pkts": r"QUIC-Tunnel packets Rx:.*\d+.*Tx:.*(\d+)",
        "forward_rx_pkts": r"QUIC-Forward packets Rx:.*(\d+).*Tx:.*\d+",
        "forward_tx_pkts": r"QUIC-Forward packets Rx:.*\d+.*Tx:.*(\d+)",
    }
    
    # Clean ANSI escape sequences from output
    stdout = ansi_escape.sub("", stdout)
    resp = dict.fromkeys(headers, 0)
    
    # Extract counters using regex patterns
    for counter, regexp in counter_regexp.items():
        counter_value = get_counter(regexp, stdout)
        if isinstance(counter_value, str):
            counter_value = get_value(counter_value)
        resp[counter] = counter_value
    
    # Determine output format and adjust headers accordingly
    quic_info_regexp = r"\d+\.\d+\w*\s+\d+\.\d+\w*\s+\d+\.\d+\w*\s+\d+\.\d+\w*\s+\d+\s+\d+\s+\d+\s+\d+\s+\d+\s+\d+"
    
    if not "INNER CONNECTIONS" in stdout:
        # Simplified format (older versions)
        quic_info_regexp = r"\d+\.\d+\w*\s+\d+\.\d+\w*\s+\d+\.\d+\w*\s+\d+\.\d+\w*\s+\d+\s+\d+\s+\d+\s+\d+"
        headers[4] = "quic_connections_created"
        headers[5] = "quic_connections_closed"
        headers = headers[:-4]  # Remove unused headers
    else:
        # Full format (newer versions)
        headers = headers[:-2]
    
    # Extract QUIC information line
    quic_info = next(
        (
            line.split()
            for line in reversed(stdout.split("\n"))
            if re.search(quic_info_regexp, line)
        ),
        None,
    )
    
    # Parse QUIC info if found
    if quic_info:
        for i, header in enumerate(headers):
            if i < len(quic_info):
                resp[header] = get_value(quic_info[i])
    
    return resp


def run_task_attempt_with_timeout_fixed(attempt_id, arguments, is_retry=False, timeout=COMMAND_TIMEOUT):
    """
    Execute a single task attempt with proper timeout handling.
    
    This function manages the complete execution lifecycle of a command:
    1. Argument preparation
    2. Command execution with timeout
    3. Log saving
    4. Output processing
    5. Status determination
    
    Args:
        attempt_id (str): Unique identifier for this attempt
        arguments (list): Command arguments
        is_retry (bool): Whether this is a retry attempt
        timeout (int): Maximum execution time in seconds
    
    Returns:
        dict: Complete attempt data including status, metrics, and logs
    """
    attempt_data = {
        "attempt_id": attempt_id,
        "status": "running",
        "arguments": arguments.copy(),
        "start_time": datetime.now().isoformat(),
        "is_retry": is_retry,
    }
    
    try:
        # Add IPv4 pool if not specified in arguments
        if "ipv4-pool" not in " ".join(arguments):
            ip_pool = re.findall(
                r"inet (23\.\d+\.\d+\.\d+)", os.popen("ip add show").read()
            )
            if ip_pool:
                arguments.append(f"--ipv4-pool={ip_pool[0]}/24")
        
        # Construct the complete command
        command = f"{HTTP3_CLIENT} {' '.join(arguments)}"
        attempt_data["command"] = command
        
        print(f"Executing attempt {attempt_id}: {command}")
        
        # Execute command with timeout
        start_time = time.time()
        result = safe_subprocess_run(command, timeout)
        execution_time = time.time() - start_time
        
        attempt_data["end_time"] = datetime.now().isoformat()
        attempt_data["execution_time"] = execution_time
        
        # Prepare metadata for logging
        metadata = {
            "attempt_id": attempt_id,
            "command": command,
            "start_time": attempt_data["start_time"],
            "end_time": attempt_data["end_time"],
            "execution_time": execution_time,
            "timeout": result['timeout'],
            "returncode": result['returncode'],
        }
        
        # Save execution logs
        log_file = save_task_log(
            attempt_id=attempt_id,
            command=command,
            stdout=result['stdout'],
            stderr=result['stderr'],
            metadata=metadata
        )
        attempt_data["log_file"] = log_file
        
        # Determine attempt status based on execution results
        if result['timeout']:
            attempt_data["status"] = "timeout"
            attempt_data["error"] = f"Command timeout after {timeout} seconds"
        elif result['returncode'] == 0:
            attempt_data["status"] = "completed"
        else:
            attempt_data["status"] = "failed"
            attempt_data["error"] = result['stderr'][-500:]  # Last 500 chars of stderr
        
        # Process output to extract metrics
        combined_output = result['stderr'] + result['stdout']
        injection_summary = process_output(combined_output)
        
        attempt_data["stats"] = injection_summary
        attempt_data["returncode"] = result['returncode']
        attempt_data["output"] = result['stdout'][-500:] if len(result['stdout']) > 500 else result['stdout']
        
        print(f"Attempt {attempt_id} completed: {attempt_data['status']}, time: {execution_time:.2f}s")
        
        return attempt_data
        
    except Exception as e:
        # Handle unexpected errors during attempt execution
        import traceback
        error_msg = str(e)
        traceback_str = traceback.format_exc()
        
        attempt_data["status"] = "task failed"
        attempt_data["error"] = error_msg
        attempt_data["error_details"] = traceback_str
        attempt_data["end_time"] = datetime.now().isoformat()
        attempt_data["execution_time"] = (
            datetime.fromisoformat(attempt_data["end_time"]) - 
            datetime.fromisoformat(attempt_data["start_time"])
        ).total_seconds()
        
        # Save error logs
        try:
            log_file = save_task_log(
                attempt_id=attempt_id,
                command=attempt_data.get("command", "Unknown"),
                stdout="",
                stderr=f"Error: {error_msg}\n{traceback_str}",
                metadata=attempt_data
            )
            attempt_data["log_file"] = log_file
        except Exception as log_error:
            print(f"Failed to save error log for attempt {attempt_id}: {log_error}")
        
        print(f"Error in attempt {attempt_id}: {error_msg}")
        return attempt_data


def run_task_with_retry_fixed(task_id, arguments):
    """
    Execute a task with automatic retry logic for duration-based tasks.
    
    This function implements the retry strategy:
    1. Extracts duration from arguments
    2. Executes attempts until duration is met or exceeded
    3. Adjusts remaining time for each retry
    4. Implements adaptive waiting between attempts
    
    Args:
        task_id (str): Unique identifier for the task
        arguments (list): Command arguments
    """
    original_arguments = arguments.copy()
    duration = extract_duration_from_args(arguments)
    
    # Update task status
    tasks[task_id]["status"] = "running"
    tasks[task_id]["start_time"] = datetime.now().isoformat()
    
    print(f"Starting task {task_id} with duration={duration}s")
    
    # Handle tasks without duration (single attempt)
    if duration is None:
        attempt_id = create_attempt_id(task_id, 1)
        attempt_data = run_task_attempt_with_timeout_fixed(attempt_id, arguments)
        
        tasks[task_id]["attempts"] = [attempt_data]
        tasks[task_id]["status"] = attempt_data["status"]
        tasks[task_id]["last_attempt"] = attempt_data
        tasks[task_id]["total_attempts"] = 1
        tasks[task_id]["total_execution_time"] = attempt_data.get("execution_time", 0)
        tasks[task_id]["has_duration"] = False
        
        save_task_metadata(task_id, tasks[task_id])
        return
    
    # Duration-based task execution
    start_time = datetime.now()
    expected_end_time = start_time + timedelta(seconds=duration)
    
    retry_count = 0
    total_execution_time = 0
    attempts = []
    
    # Initialize task duration information
    tasks[task_id]["has_duration"] = True
    tasks[task_id]["duration_seconds"] = duration
    tasks[task_id]["requested_duration"] = duration
    
    print(f"Task {task_id} should run until: {expected_end_time}")
    
    # Main retry loop
    while datetime.now() < expected_end_time:
        retry_count += 1
        attempt_id = create_attempt_id(task_id, retry_count)
        
        # Calculate remaining time
        remaining_time = max(1, (expected_end_time - datetime.now()).total_seconds())
        print(f"Attempt {retry_count}, remaining time: {remaining_time:.0f}s")
        
        # Update duration argument for this attempt
        current_arguments = original_arguments.copy()
        for i, arg in enumerate(current_arguments):
            if arg.startswith("--duration="):
                arg_parts = arg.split()
                if len(arg_parts) > 1:
                    # Handle arguments with appended parameters
                    new_args = []
                    for part in arg_parts:
                        if part.startswith("--duration="):
                            new_args.append(f"--duration={remaining_time:.0f}")
                        else:
                            new_args.append(part)
                    current_arguments[i:i+1] = new_args
                else:
                    current_arguments[i] = f"--duration={remaining_time:.0f}"
            elif arg == "--duration" and i + 1 < len(current_arguments):
                current_arguments[i + 1] = f"{remaining_time:.0f}"
        
        # Calculate timeout with safety margin
        attempt_timeout = min(remaining_time + 300, COMMAND_TIMEOUT)  # +5 minutes margin
        is_retry = retry_count > 1
        
        # Execute attempt
        attempt_data = run_task_attempt_with_timeout_fixed(
            attempt_id, current_arguments, is_retry=is_retry, timeout=attempt_timeout
        )
        
        attempts.append(attempt_data)
        exec_time = attempt_data.get("execution_time", 0)
        total_execution_time += exec_time
        
        # Update task progress
        tasks[task_id]["attempts"] = attempts
        tasks[task_id]["total_attempts"] = len(attempts)
        tasks[task_id]["successful_attempts"] = sum(
            1 for a in attempts if a.get("status") == "completed"
        )
        tasks[task_id]["failed_attempts"] = sum(
            1 for a in attempts if a.get("status") in ["failed", "task failed", "timeout"]
        )
        tasks[task_id]["last_attempt"] = attempt_data
        tasks[task_id]["total_execution_time"] = total_execution_time
        tasks[task_id]["achieved_duration_ratio"] = (
            total_execution_time / duration if duration > 0 else 1.0
        )
        
        print(f"Attempt {retry_count} completed. Status: {attempt_data['status']}, "
              f"Time: {exec_time:.2f}s, Total: {total_execution_time:.2f}s")
        
        # Check completion criteria
        if attempt_data.get("status") == "completed" and total_execution_time >= duration:
            print(f"Task {task_id} completed requested duration")
            break
        
        if datetime.now() >= expected_end_time:
            print(f"Task {task_id} reached time limit")
            break
        
        # Adaptive waiting between attempts
        if exec_time < 30:  # If attempt failed quickly
            wait_time = min(30, remaining_time / 2)
            print(f"Waiting {wait_time:.1f}s before next attempt...")
            time.sleep(wait_time)
    
    # Calculate final statistics
    successful_attempts = sum(1 for a in attempts if a.get("status") == "completed")
    failed_attempts = sum(
        1 for a in attempts if a.get("status") in ["failed", "task failed", "timeout"]
    )
    
    # Finalize task data
    tasks[task_id]["attempts"] = attempts
    tasks[task_id]["total_attempts"] = len(attempts)
    tasks[task_id]["successful_attempts"] = successful_attempts
    tasks[task_id]["failed_attempts"] = failed_attempts
    tasks[task_id]["last_attempt"] = attempts[-1] if attempts else None
    tasks[task_id]["total_execution_time"] = total_execution_time
    tasks[task_id]["start_time"] = start_time.isoformat()
    tasks[task_id]["end_time"] = datetime.now().isoformat()
    
    # Determine final task status
    if total_execution_time >= duration:
        tasks[task_id]["status"] = "completed_with_retries"
    elif failed_attempts == len(attempts):
        tasks[task_id]["status"] = "failed_all_attempts"
    else:
        tasks[task_id]["status"] = "partial_completion"
    
    # Save final metadata
    save_task_metadata(task_id, tasks[task_id])
    
    print(f"Task {task_id} finished with status: {tasks[task_id]['status']}")


def run_single_attempt_wrapper(task_id, arguments, duration=None):
    """
    Wrapper function for executing a single attempt in a separate thread.
    
    This function is used for tasks without auto-retry to ensure
    non-blocking execution while maintaining proper error handling.
    
    Args:
        task_id (str): Unique identifier for the task
        arguments (list): Command arguments
        duration (int, optional): Expected duration in seconds
    """
    try:
        tasks[task_id]["status"] = "running"
        tasks[task_id]["start_time"] = datetime.now().isoformat()
        
        attempt_id = create_attempt_id(task_id, 1)
        attempt_data = run_task_attempt_with_timeout_fixed(
            attempt_id, arguments, is_retry=False
        )
        
        tasks[task_id]["attempts"] = [attempt_data]
        tasks[task_id]["status"] = attempt_data["status"]
        tasks[task_id]["last_attempt"] = attempt_data
        tasks[task_id]["total_attempts"] = 1
        tasks[task_id]["total_execution_time"] = attempt_data.get("execution_time", 0)
        
        if duration:
            tasks[task_id]["has_duration"] = True
            tasks[task_id]["duration_seconds"] = duration
            tasks[task_id]["requested_duration"] = duration
            
            # Check if duration requirement was met
            if attempt_data.get("status") == "completed":
                if attempt_data.get("execution_time", 0) >= duration:
                    tasks[task_id]["status"] = "completed"
                else:
                    tasks[task_id]["status"] = "failed_before_duration"
        
        save_task_metadata(task_id, tasks[task_id])
        
    except Exception as e:
        print(f"Error in run_single_attempt_wrapper for task {task_id}: {e}")
        tasks[task_id]["status"] = "task failed"
        tasks[task_id]["error"] = str(e)
        tasks[task_id]["end_time"] = datetime.now().isoformat()
        
        try:
            save_task_metadata(task_id, tasks[task_id])
        except Exception as meta_error:
            print(f"Failed to save metadata for failed task {task_id}: {meta_error}")


# =====================================================================
# FLASK API ENDPOINTS
# =====================================================================

@app.route("/tasks", methods=["POST"])
def create_task():
    """
    Create and start a new HTTP3 client task.
    
    This endpoint accepts task configuration and immediately returns
    a task ID while starting execution in a background thread.
    
    Request JSON:
    {
        "arguments": ["--duration=300", "--target=example.com"],
        "auto_retry": true  # Optional, defaults to true
    }
    
    Returns:
        JSON response with 202 Accepted status containing task information
    """
    arguments = request.json.get("arguments")
    if not arguments or not isinstance(arguments, list):
        return jsonify({"error": "Invalid arguments format"}), 400
    
    auto_retry = request.json.get("auto_retry", True)
    task_id = str(uuid.uuid4())
    
    # Initialize task structure
    tasks[task_id] = {
        "task_id": task_id,
        "status": "initializing",
        "arguments": arguments.copy(),
        "auto_retry_enabled": auto_retry,
        "created_at": datetime.now().isoformat(),
        "attempts": [],
    }
    
    duration = extract_duration_from_args(arguments)
    print(f"Creating task {task_id} with duration={duration}, auto_retry={auto_retry}")
    
    # Always execute in a separate thread for non-blocking response
    if duration and auto_retry:
        # Use retry logic for duration-based tasks with auto-retry enabled
        tasks[task_id]["has_duration"] = True
        tasks[task_id]["duration_seconds"] = duration
        tasks[task_id]["requested_duration"] = duration
        
        thread = threading.Thread(
            target=run_task_with_retry_fixed,
            args=(task_id, arguments),
            daemon=True,
        )
    else:
        # Single attempt execution
        tasks[task_id]["has_duration"] = duration is not None
        if duration:
            tasks[task_id]["duration_seconds"] = duration
            tasks[task_id]["requested_duration"] = duration
        
        thread = threading.Thread(
            target=run_single_attempt_wrapper,
            args=(task_id, arguments, duration),
            daemon=True,
        )
    
    # Start execution and track thread
    task_threads[task_id] = thread
    thread.start()
    
    # Immediate response with task information
    return jsonify({
        "task_id": task_id,
        "status": "initializing",
        "message": "Task created and started",
        "duration_detected": duration is not None,
        "auto_retry": auto_retry,
        "has_duration": duration is not None,
        "duration_seconds": duration,
        "created_at": tasks[task_id]["created_at"]
    }), 202  # 202 Accepted - task started but not yet completed


@app.route("/tasks/<task_id>", methods=["GET", "DELETE"])
def get_task(task_id):
    """
    Retrieve or delete a specific task.
    
    GET: Returns comprehensive task information including execution status,
         attempts, and performance metrics.
    
    DELETE: Terminates the task and cleans up associated resources.
    
    Args:
        task_id (str): Unique identifier of the task
    
    Returns:
        GET: JSON with task details
        DELETE: JSON with confirmation message
    """
    task = tasks.get(task_id)
    if not task:
        return jsonify({"error": "Task not found"}), 404
    
    if request.method == "GET":
        # Build comprehensive task response
        response_data = {
            "task_id": task["task_id"],
            "status": task["status"],
            "arguments": task["arguments"],
            "created_at": task["created_at"],
            "total_attempts": task.get("total_attempts", 0),
            "successful_attempts": task.get("successful_attempts", 0),
            "failed_attempts": task.get("failed_attempts", 0),
            "total_execution_time": task.get("total_execution_time", 0),
            "requested_duration": task.get("requested_duration"),
            "achieved_duration_ratio": task.get("achieved_duration_ratio"),
            "has_duration": task.get("has_duration", False),
            "auto_retry_enabled": task.get("auto_retry_enabled", False),
            "last_attempt": task.get("last_attempt"),
            "attempts_summary": [
                {
                    "attempt_id": a.get("attempt_id"),
                    "status": a.get("status"),
                    "execution_time": a.get("execution_time"),
                    "returncode": a.get("returncode"),
                    "start_time": a.get("start_time"),
                    "end_time": a.get("end_time"),
                }
                for a in task.get("attempts", [])
            ]
            if task.get("attempts")
            else [],
        }
        
        # Include stats if only one attempt exists
        if len(task.get("attempts", [])) == 1:
            response_data["stats"] = task["attempts"][0].get("stats", {})
        
        return jsonify(response_data)
    
    elif request.method == "DELETE":
        terminated_processes = []
        
        # Terminate all associated subprocesses
        for attempt in task.get("attempts", []):
            attempt_id = attempt.get("attempt_id")
            process = processes.pop(attempt_id, None)
            if process and process.poll() is None:
                try:
                    # Kill entire process group
                    os.killpg(os.getpgid(process.pid), signal.SIGTERM)
                    terminated_processes.append(attempt_id)
                except (OSError, ProcessLookupError) as e:
                    print(f"Error terminating process for attempt {attempt_id}: {e}")
                    continue
        
        # Update task status
        tasks[task_id]["status"] = "terminated"
        tasks[task_id]["end_time"] = datetime.now().isoformat()
        tasks[task_id]["terminated_processes"] = terminated_processes
        
        # Clean up thread tracking
        if task_id in task_threads:
            del task_threads[task_id]
        
        return jsonify({
            "message": f"Task {task_id} successfully terminated",
            "terminated_processes": terminated_processes,
        }), 200


@app.route("/tasks", methods=["GET"])
def get_all_tasks():
    """
    Retrieve information about all tasks.
    
    Supports optional date filtering via query parameters:
    - start_date: Unix timestamp for start of range
    - end_date: Unix timestamp for end of range
    
    Query Parameters:
        start_date (int, optional): Filter tasks starting after this timestamp
        end_date (int, optional): Filter tasks ending before this timestamp
    
    Returns:
        JSON array of task summaries
    """
    start_date = request.args.get("start_date", 0)
    end_date = request.args.get("end_date", (1 << 63) - 1)  # Maximum 64-bit integer

    if start_date == "0" and end_date == str((1 << 63) - 1):
        # Return all tasks
        resp = []
        for task_id, task_data in tasks.items():
            row = {
                "task_id": task_id,
                "status": task_data.get("status", "unknown"),
                "arguments": task_data.get("arguments", []),
                "created_at": task_data.get("created_at"),
                "start_time": task_data.get("start_time"),
                "end_time": task_data.get("end_time"),
                "total_attempts": task_data.get("total_attempts", 0),
                "has_duration": task_data.get("has_duration", False),
                "total_execution_time": task_data.get("total_execution_time", 0),
            }

            # Extract output and error information
            if task_data.get("attempts"):
                last_attempt = task_data["attempts"][-1]
                row["output"] = ansi_escape.sub("", last_attempt.get("output", ""))
                error = last_attempt.get("error", "")
                error = ansi_escape.sub("", error)
                row["error"] = "\n".join(error.split("\n")[-4:-1])  # Last 4 lines
            else:
                row["output"] = ansi_escape.sub("", task_data.get("output", ""))
                error = task_data.get("error", "")
                error = ansi_escape.sub("", error)
                row["error"] = "\n".join(error.split("\n")[-4:-1])

            resp.append(row)
        return jsonify(resp)
    else:
        # Date-filtered response
        resp = []
        for task_data in filter_tasks_by_date(
            start_time=int(start_date), end_time=int(end_date)
        ):
            row = task_data
            stderr = task_data.get("error", "")
            stderr = ansi_escape.sub("", stderr)
            row["error"] = "\n".join(stderr.split("\n")[-4:-1])
            row["output"] = ansi_escape.sub("", task_data.get("output", ""))
            resp.append(row)
        return jsonify(resp)


@app.route("/tasks/summary", methods=["GET"])
def get_summary():
    """
    Generate a statistical summary of all tasks.
    
    Provides aggregated statistics including:
    - Task counts by status
    - Attempt statistics
    - Performance metrics totals
    - Average execution metrics
    
    Query Parameters:
        start_date (int, optional): Filter tasks starting after this timestamp
        end_date (int, optional): Filter tasks ending before this timestamp
    
    Returns:
        JSON object with comprehensive summary statistics
    """
    start_date = request.args.get("start_date", 0)
    end_date = request.args.get("end_date", (1 << 63) - 1)
    filtered_tasks = list(
        filter_tasks_by_date(start_time=int(start_date), end_time=int(end_date))
    )

    # Initialize summary structure
    summary = {
        # Task status counts
        "total": len(tasks),
        "running": sum(1 for task in tasks.values() if task.get("status") == "running"),
        "initializing": sum(
            1 for task in tasks.values() if task.get("status") == "initializing"
        ),
        "completed": sum(
            1 for task in filtered_tasks if task.get("status") == "completed"
        ),
        "failed": sum(1 for task in filtered_tasks if task.get("status") == "failed"),
        "task failed": sum(
            1 for task in filtered_tasks if task.get("status") == "task failed"
        ),
        "completed_with_retries": sum(
            1
            for task in filtered_tasks
            if task.get("status") == "completed_with_retries"
        ),
        "partial_completion": sum(
            1 for task in filtered_tasks if task.get("status") == "partial_completion"
        ),
        "failed_all_attempts": sum(
            1 for task in filtered_tasks if task.get("status") == "failed_all_attempts"
        ),
        "failed_before_duration": sum(
            1
            for task in filtered_tasks
            if task.get("status") == "failed_before_duration"
        ),
        "terminated": sum(
            1 for task in filtered_tasks if task.get("status") == "terminated"
        ),
        "timeout": sum(1 for task in filtered_tasks if task.get("status") == "timeout"),
        
        # Attempt statistics
        "total_attempts": 0,
        "successful_attempts": 0,
        "failed_attempts": 0,
        "total_execution_time": 0,
        "average_attempts_per_task": 0,
    }

    # Performance metric names
    stat_names = [
        "http3_errors",
        "masque_errors",
        "tunnel_tx_pkts",
        "tunnel_rx_pkts",
        "forward_tx_pkts",
        "forward_rx_pkts",
        "inner_connections_closed",
        "inner_connections_created",
        "outer_connections_closed",
        "outer_connections_created",
        "quic_streams_closed",
        "quic_streams_created",
        "udp_bytes_tx",
        "udp_bytes_rx",
        "payload_bytes",
        "total_time",
    ]

    # Initialize metric accumulators
    for stat_name in stat_names:
        summary[f"total_{stat_name}"] = 0

    total_attempts_all = 0
    successful_attempts_all = 0
    failed_attempts_all = 0
    total_execution_time_all = 0

    # Aggregate statistics from all filtered tasks
    for task in filtered_tasks:
        total_attempts_all += task.get("total_attempts", 1)
        successful_attempts_all += task.get("successful_attempts", 0)
        failed_attempts_all += task.get("failed_attempts", 0)
        total_execution_time_all += task.get("total_execution_time", 0)

        # Aggregate performance metrics
        attempts = task.get("attempts", [])
        if attempts:
            for attempt in attempts:
                stats = attempt.get("stats", {})
                for stat_name in stat_names:
                    summary[f"total_{stat_name}"] += stats.get(stat_name, 0)
        else:
            # Legacy support for tasks without attempts array
            stats = task.get("stats", {})
            for stat_name in stat_names:
                summary[f"total_{stat_name}"] += stats.get(stat_name, 0)

    # Update summary with aggregated statistics
    summary["total_attempts"] = total_attempts_all
    summary["successful_attempts"] = successful_attempts_all
    summary["failed_attempts"] = failed_attempts_all
    summary["total_execution_time"] = total_execution_time_all

    # Calculate averages
    if len(filtered_tasks) > 0:
        summary["average_attempts_per_task"] = total_attempts_all / len(filtered_tasks)

    return jsonify(summary)


# =====================================================================
# LOG MANAGEMENT ENDPOINTS
# =====================================================================

@app.route("/logs", methods=["GET"])
def get_all_logs():
    """
    Download all logs as a compressed tarball (.tgz).
    
    This endpoint creates an on-demand archive of all log files
    and streams it to the client for download.
    
    Returns:
        Binary tarball file with 200 OK, or error message with 404/500
    """
    try:
        tarball_path, error = create_logs_tarball()
        
        if error:
            return jsonify({"error": error}), 404
        
        filename = os.path.basename(tarball_path)
        
        # Stream file to client
        response = send_file(
            tarball_path,
            as_attachment=True,
            download_name=filename,
            mimetype='application/gzip'
        )
        
        # Clean up temporary files after sending
        @response.call_on_close
        def cleanup():
            try:
                if os.path.exists(tarball_path):
                    os.remove(tarball_path)
                    temp_dir = os.path.dirname(tarball_path)
                    if os.path.exists(temp_dir):
                        os.rmdir(temp_dir)
            except Exception as cleanup_error:
                print(f"Error cleaning up temporary files: {cleanup_error}")
        
        return response
        
    except Exception as e:
        return jsonify({"error": f"Failed to create logs archive: {str(e)}"}), 500


@app.route("/logs/<task_id>", methods=["GET"])
def get_task_logs(task_id):
    """
    Download logs for a specific task as a compressed tarball (.tgz).
    
    Args:
        task_id (str): Unique identifier of the task
    
    Returns:
        Binary tarball file with 200 OK, or error message with 404/500
    """
    try:
        tarball_path, error = create_logs_tarball(task_id=task_id)
        
        if error:
            return jsonify({"error": error}), 404
        
        filename = os.path.basename(tarball_path)
        
        response = send_file(
            tarball_path,
            as_attachment=True,
            download_name=filename,
            mimetype='application/gzip'
        )
        
        @response.call_on_close
        def cleanup():
            try:
                if os.path.exists(tarball_path):
                    os.remove(tarball_path)
                    temp_dir = os.path.dirname(tarball_path)
                    if os.path.exists(temp_dir):
                        os.rmdir(temp_dir)
            except Exception as cleanup_error:
                print(f"Error cleaning up temporary files: {cleanup_error}")
        
        return response
        
    except Exception as e:
        return jsonify({"error": f"Failed to create logs archive: {str(e)}"}), 500


@app.route("/logs", methods=["DELETE"])
def delete_all_logs():
    """
    Delete all log files from the system.
    
    Returns:
        JSON confirmation message with 200 OK, or error message with 404
    """
    success, message = delete_logs()
    
    if success:
        return jsonify({"message": message}), 200
    else:
        return jsonify({"error": message}), 404


@app.route("/logs/<task_id>", methods=["DELETE"])
def delete_task_logs(task_id):
    """
    Delete log files for a specific task.
    
    Args:
        task_id (str): Unique identifier of the task
    
    Returns:
        JSON confirmation message with 200 OK, or error message with 404
    """
    success, message = delete_logs(task_id=task_id)
    
    if success:
        return jsonify({"message": message}), 200
    else:
        return jsonify({"error": message}), 404


@app.route("/logs/info", methods=["GET"])
def get_logs_info():
    """
    Get information about stored logs.
    
    Provides metadata about log storage including:
    - Total size and file counts
    - Per-task statistics
    - Storage configuration
    
    Returns:
        JSON object with log storage information
    """
    try:
        if not os.path.exists(LOG_BASE_DIR):
            return jsonify({
                "total_size": 0,
                "total_tasks": 0,
                "total_log_files": 0,
                "log_directory": LOG_BASE_DIR,
                "retention_days": LOG_RETENTION_DAYS,
                "tasks": []
            })
        
        total_size = 0
        total_tasks = 0
        total_log_files = 0
        tasks_info = []
        
        tasks_dir = os.path.join(LOG_BASE_DIR, "tasks")
        if os.path.exists(tasks_dir):
            for task_id in os.listdir(tasks_dir):
                task_dir = os.path.join(tasks_dir, task_id)
                if os.path.isdir(task_dir):
                    task_size = 0
                    task_log_files = 0
                    
                    # Calculate directory size and file count
                    for root, dirs, files in os.walk(task_dir):
                        for file in files:
                            file_path = os.path.join(root, file)
                            task_size += os.path.getsize(file_path)
                            task_log_files += 1
                    
                    total_size += task_size
                    total_tasks += 1
                    total_log_files += task_log_files
                    
                    # Load task metadata if available
                    metadata_file = os.path.join(task_dir, "metadata.json")
                    metadata = {}
                    if os.path.exists(metadata_file):
                        try:
                            with open(metadata_file, 'r') as f:
                                metadata = json.load(f)
                        except (json.JSONDecodeError, IOError) as e:
                            print(f"Error reading metadata for task {task_id}: {e}")
                    
                    tasks_info.append({
                        "task_id": task_id,
                        "size_bytes": task_size,
                        "size_human": f"{task_size / (1024*1024):.2f} MB",
                        "log_files": task_log_files,
                        "status": metadata.get("status", "unknown"),
                        "created_at": metadata.get("created_at", "unknown"),
                        "attempts": metadata.get("total_attempts", 0)
                    })
        
        # Sort by creation date (newest first)
        tasks_info.sort(key=lambda x: x.get("created_at", ""), reverse=True)
        
        return jsonify({
            "total_size": total_size,
            "total_size_human": f"{total_size / (1024*1024):.2f} MB",
            "total_tasks": total_tasks,
            "total_log_files": total_log_files,
            "log_directory": LOG_BASE_DIR,
            "retention_days": LOG_RETENTION_DAYS,
            "tasks": tasks_info[:50]  # Limit to 50 most recent tasks
        })
        
    except Exception as e:
        return jsonify({"error": f"Failed to get logs information: {str(e)}"}), 500


# =====================================================================
# SUPPORTING FUNCTIONS
# =====================================================================

def filter_tasks_by_date(start_time=0, end_time=(1 << 63) - 1):
    """
    Filter tasks by their execution time range.
    
    Args:
        start_time (int): Unix timestamp for start of range
        end_time (int): Unix timestamp for end of range
    
    Yields:
        dict: Task data for tasks within the specified time range
    """
    for task_id, task in tasks.items():
        try:
            # For running tasks, use current time as end time
            if task.get("status") in ["initializing", "running"]:
                task_end_time = datetime.now().timestamp()
            else:
                task_end_time = datetime.fromisoformat(
                    task.get("end_time", "2036-01-01")
                ).timestamp()

            task_start_time = datetime.fromisoformat(
                task.get("start_time", task.get("created_at", "1970-01-01"))
            ).timestamp()

            if task_start_time > start_time and task_end_time < end_time:
                yield task
        except Exception as e:
            print(f"Error filtering task {task_id}: {e}")
            continue


# =====================================================================
# CERTIFICATE MANAGEMENT ENDPOINTS (Existing functionality)
# =====================================================================

def process_cert_output(stdout):
    """
    Process certificate retriever output to extract certificate information.
    
    Args:
        stdout (str): JSON output from cert-retriever command
    
    Returns:
        dict: Structured certificate information
    """
    raw_data = json.loads(stdout)
    valid_data = [i for i in raw_data if not i.get("error")]
    error_data = [i for i in raw_data if i.get("error")]
    resp = {
        "different_certs": len(valid_data),
        "certificates": [],
        "errors": error_data,
    }
    
    for cert in valid_data:
        certs = []
        for chain in cert["presented"]["certificates"]:
            sni = chain.get("subject").get("commonName")
            issuer = chain.get("issuer").get("commonName")
            valid_from = chain.get("notBefore")
            valid_until = chain.get("notAfter")
            certs.append({
                "sni": sni,
                "issuer": issuer,
                "valid_from": valid_from,
                "valid_until": valid_until,
            })
        resp["server_sni"] = cert["presented"].get("serverSNI", "")
        resp["chain_length"] = cert["presented"].get("chainLength", "1")
        resp["is_complete"] = cert["presented"].get("isComplete", False)
        resp["missing_root"] = cert["presented"].get("missingRoot", False)
        resp["certificates"].append(certs)
    
    return resp


def get_certificate(task_id, arguments, protocol="udp"):
    """
    Execute certificate retrieval command and process results.
    
    Args:
        task_id (str): Unique identifier for the certificate task
        arguments (dict): Certificate retrieval arguments
        protocol (str): Protocol to use (udp/tcp)
    """
    stdout = b""
    stderr = b""
    try:
        # Add IPv4 pool if not specified
        if "ipv4-pool" not in " ".join(arguments):
            ip_pool = re.findall(
                r"inet (23\.\d+\.\d+\.\d+)", os.popen("ip add show").read()
            )
            if ip_pool:
                arguments["ipv4-pool"] = f"{ip_pool[0]}/24"
        
        # Build command string
        argos = ""
        for k, v in arguments.items():
            argos += f" -{k} {v}"
        command = f"sleep 10; {CERTIFICATE_RETRIEVER} {argos}"
        
        # Update task information
        cert_tasks[task_id]["command"] = command
        process = subprocess.Popen(
            command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE
        )
        processes[task_id] = process
        cert_tasks[task_id]["start_time"] = datetime.now().isoformat()
        
        # Execute and capture output
        stdout, stderr = process.communicate()
        cert_tasks[task_id]["end_time"] = datetime.now().isoformat()
        
        stdout_decoded = stdout.decode("utf-8")
        stderr_decoded = stderr.decode("utf-8")
        
        # Process certificate data
        certificates_summary = process_cert_output(stderr_decoded + stdout_decoded)
        
        # Determine task status
        task_result = "completed"
        if process.returncode != 0:
            task_result = "failed"

        cert_tasks[task_id]["status"] = task_result
        cert_tasks[task_id]["output"] = stdout_decoded
        cert_tasks[task_id]["error"] = stderr_decoded
        cert_tasks[task_id]["certificates_data"] = certificates_summary
        
    except Exception as e:
        cert_tasks[task_id]["status"] = "task failed"
        cert_tasks[task_id]["error"] = str(e)
        cert_tasks[task_id]["output"] = (
            f"stderr:\n{stderr.decode()}\nstdout:\n{stdout.decode()}"
        )


@app.route("/certificates", methods=["POST"])
def certificates():
    """
    Create a new certificate retrieval task.
    
    Request JSON:
    {
        "arguments": {
            "server": "example.com",
            "port": 443
        }
    }
    
    Returns:
        JSON with task ID and 202 Accepted status
    """
    arguments = request.json.get("arguments")
    if not arguments or not isinstance(arguments, dict):
        return jsonify({"error": "Invalid arguments format"}), 400

    task_id = str(uuid.uuid4())
    cert_tasks[task_id] = {
        "status": "running",
        "arguments": arguments,
        "task_id": task_id,
    }
    
    # Execute in background thread
    thread = threading.Thread(target=get_certificate, args=(task_id, arguments))
    thread.start()
    
    return jsonify({"task_id": task_id}), 202


@app.route("/certificates", methods=["GET"])
def retrieve_certificates():
    """
    Retrieve information about all certificate tasks.
    
    Returns:
        JSON object containing all certificate tasks
    """
    return jsonify(cert_tasks)


@app.route("/certificates/last", methods=["GET"])
def retrieve_last_certificates():
    """
    Retrieve the most recent certificate task.
    
    Returns:
        JSON object with the latest certificate task data
    """
    last_cert = (
        max(cert_tasks.values(), key=lambda x: datetime.fromisoformat(x["start_time"]))
        if cert_tasks
        else {}
    )
    return jsonify(last_cert)


@app.route("/certificates/<task_id>", methods=["GET"])
def retrieve_certificate(task_id):
    """
    Retrieve a specific certificate task.
    
    Args:
        task_id (str): Unique identifier of the certificate task
    
    Returns:
        JSON object with certificate task data
    """
    resp = cert_tasks.get(task_id, {})
    return jsonify(resp)


# =====================================================================
# DOCUMENTATION ENDPOINTS
# =====================================================================

@app.route("/openapi.json")
def openapi():
    """
    Serve OpenAPI specification file.
    
    Returns:
        OpenAPI JSON specification
    """
    return send_from_directory(".", "openapi-client.json")


@app.route("/docs")
def docs():
    """
    Serve API documentation interface.
    
    Returns:
        HTML documentation interface
    """
    return send_from_directory(".", "index.html")


# =====================================================================
# APPLICATION INITIALIZATION
# =====================================================================

if __name__ == "__main__":
    """
    Main application entry point.
    
    Initializes the application with:
    1. Process pool for concurrent execution
    2. Log cleanup scheduler
    3. Flask server configuration
    """
    # Initialize process pool
    init_process_pool()
    
    # Initial log cleanup
    cleanup_thread = threading.Thread(target=cleanup_old_logs, daemon=True)
    cleanup_thread.start()
    
    # Periodic log cleanup scheduler (every 6 hours)
    def schedule_cleanup():
        while True:
            time.sleep(6 * 3600)  # 6 hours
            cleanup_old_logs()
    
    scheduler_thread = threading.Thread(target=schedule_cleanup, daemon=True)
    scheduler_thread.start()
    
    # Application startup message
    print(f"Logs directory: {LOG_BASE_DIR}")
    print(f"Log retention: {LOG_RETENTION_DAYS} days")
    print("Starting Flask server on port 80...")
    
    # Start Flask application
    app.run(
        debug=False,        # Disable debug mode for production
        port=80,            # Listen on standard HTTP port
        host="0.0.0.0",     # Listen on all network interfaces
        threaded=True,      # Enable threading for concurrent requests
    )
