package com.onebooking.onepagecommerce.entity;

import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;

import java.time.LocalDateTime;

@Entity
@Table(name = "transactions")
@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class Transaction {

    @Id
    @Column(length = 36, updatable = false)
    private String id;

    @Column(name = "buyer_id", length = 36, nullable = false)
    private String buyerId;

    @Column(name = "seller_id", length = 36, nullable = false)
    private String sellerId;

    @Column(name = "item_id", length = 36)
    private String itemId;

    @Column(nullable = false, length = 40)
    private String type;

    @Column(nullable = false)
    private double amount;

    @Column(nullable = false, length = 3)
    private String currency = "USD";

    @Column(nullable = false, length = 20)
    private String status = "PENDING";

    @Column(name = "payment_method", nullable = false, length = 40)
    private String paymentMethod;

    @Column(name = "escrow_id", length = 64)
    private String escrowId;

    @Column(name = "released_at")
    private LocalDateTime releasedAt;

    @Column(name = "completed_at")
    private LocalDateTime completedAt;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @LastModifiedDate
    @Column(nullable = false)
    private LocalDateTime updatedAt;
}
