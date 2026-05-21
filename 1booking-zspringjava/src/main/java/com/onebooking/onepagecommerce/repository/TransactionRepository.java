package com.onebooking.onepagecommerce.repository;

import com.onebooking.onepagecommerce.entity.Transaction;
import org.springframework.data.domain.*;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.*;

public interface TransactionRepository extends JpaRepository<Transaction, String> {
    List<Transaction> findByBuyerId(String buyerId);
    List<Transaction> findBySellerId(String sellerId);
    List<Transaction> findByEscrowId(String escrowId);
    Page<Transaction> findByStatus(String status, Pageable pageable);
}
