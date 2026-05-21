package com.onebooking.onepagecommerce.util;

import com.fasterxml.jackson.databind.ObjectMapper;

import java.util.Map;

public final class Helper {

    private static final ObjectMapper OBJECT_MAPPER = new ObjectMapper();

    private Helper() {}

    public static Map<String, Object> successResponse(Object data) {
        return successResponse(data, null);
    }

    public static Map<String, Object> successResponse(Object data, Object meta) {
        Map<String, Object> body = new java.util.LinkedHashMap<>();
        body.put("success", true);
        body.put("data", data);
        if (meta != null) body.put("meta", meta);
        body.put("timestamp", java.time.Instant.now().toString());
        return body;
    }

    public static Map<String, Object> errorResponse(String code, String message, Object details) {
        Map<String, Object> body = new java.util.LinkedHashMap<>();
        body.put("success", false);
        Map<String, Object> error = new java.util.LinkedHashMap<>();
        error.put("code", code);
        error.put("message", message);
        if (details != null) error.put("details", details);
        body.put("error", error);
        body.put("timestamp", java.time.Instant.now().toString());
        return body;
    }

    public static String generateUuid() {
        return java.util.UUID.randomUUID().toString().replace("-", "");
    }
}
