package com.onebooking.onepagecommerce.entity;

import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;

import java.time.LocalDate;
import java.time.LocalDateTime;

@Entity
@Table(name = "scholarships")
@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class Scholarship {

    @Id
    @Column(length = 36, updatable = false)
    private String id;

    @Column(nullable = false, length = 200)
    private String title;

    @Column(nullable = false, length = 5000)
    private String description;

    @Column(nullable = false, length = 120)
    private String provider;

    @Column(nullable = false, length = 30)
    private String providerType;

    @Column(nullable = false, length = 80)
    private String country;

    @Column(length = 80)
    private String city;

    @Column(nullable = false, length = 30)
    private String level = "UNDERGRADUATE";

    @Column(nullable = false, length = 120)
    private String field;

    @Column(nullable = false)
    private double amount = 0.0;

    @Column(nullable = false, length = 3)
    private String currency = "USD";

    @Column(nullable = false, length = -1)
    private String covers;

    private LocalDate deadline;

    @Column(length = 4000)
    private String eligibility;

    @Column(length = 4000)
    private String requirements;

    @Column(name = "application_url")
    private String applicationUrl;

    @Column(nullable = false)
    private boolean isActive = true;

    @Column(nullable = false)
    private boolean isFeatured = false;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @LastModifiedDate
    @Column(nullable = false)
    private LocalDateTime updatedAt;
}
