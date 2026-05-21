package com.onebooking.onepagecommerce.entity;

import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;

import java.time.LocalDateTime;

@Entity
@Table(name = "marketplace_items")
@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class MarketplaceItem {

    @Id
    @Column(length = 36, updatable = false)
    private String id;

    @Column(name = "seller_id", length = 36, nullable = false)
    private String sellerId;

    @Column(nullable = false, length = 120)
    private String title;

    @Column(length = 4000)
    private String description;

    @Column(nullable = false, length = 80)
    private String category;

    @Column(name = "subcategory", length = 80)
    private String subcategory;

    @Column(nullable = false)
    private double price;

    @Column(nullable = false, length = 3)
    private String currency = "USD";

    @Column(nullable = false, length = 20)
    private String condition;

    @Column(length = 120)
    private String location;

    @Column(length = 80)
    private String country;

    @Column(name = "shipping_available", nullable = false)
    private boolean shippingAvailable = false;

    @Column(name = "shipping_cost", nullable = false)
    private double shippingCost = 0.0;

    @Column(name = "image_urls", length = -1)
    private String imageUrls;

    @Column(nullable = false)
    private boolean isActive = true;

    @Column(nullable = false)
    private boolean isSold = false;

    @Column(nullable = false)
    private int viewCount = 0;

    @Column(name = "interest_count", nullable = false)
    private int interestCount = 0;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @LastModifiedDate
    @Column(nullable = false)
    private LocalDateTime updatedAt;
}
