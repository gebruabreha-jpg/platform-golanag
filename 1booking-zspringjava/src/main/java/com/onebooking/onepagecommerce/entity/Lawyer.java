package com.onebooking.onepagecommerce.entity;

import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;

import java.time.LocalDateTime;

@Entity
@Table(name = "lawyers")
@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class Lawyer {

    @Id
    @Column(length = 36, updatable = false)
    private String id;

    @Column(name = "user_id", length = 36, nullable = false, unique = true)
    private String userId;

    @Column(nullable = false, length = 80)
    private String specialization;

    @Column(nullable = false, length = 120)
    private String name;

    @Column(length = 120)
    private String firm;

    @Column(name = "years_experience")
    private Integer yearsExperience;

    @Column(length = 120)
    private String location;

    @Column(length = 80)
    private String country;

    @Column(name = "consultation_fee", nullable = false)
    private double consultationFee;

    @Column(nullable = false, length = 3)
    private String currency = "USD";

    @Column(nullable = false)
    private boolean isVerified = false;

    @Column(nullable = false)
    private double rating = 0.0;

    @Column(nullable = false)
    private int reviewCount = 0;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @LastModifiedDate
    @Column(nullable = false)
    private LocalDateTime updatedAt;
}
