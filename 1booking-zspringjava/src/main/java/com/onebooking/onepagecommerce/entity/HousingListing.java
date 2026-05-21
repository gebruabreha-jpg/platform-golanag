package com.onebooking.onepagecommerce.entity;

import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;

import java.time.LocalDateTime;

@Entity
@Table(name = "housing_listings")
@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class HousingListing {

    @Id
    @Column(length = 36, updatable = false)
    private String id;

    @Column(name = "landlord_id", length = 36, nullable = false)
    private String landlordId;

    @Column(nullable = false, length = 120)
    private String title;

    @Column(length = 4000)
    private String description;

    @Column(nullable = false, length = 30)
    private String propertyType;

    @Column(nullable = false, length = 20)
    private String roomType;

    @Column(nullable = false)
    private int bedrooms;

    @Column(nullable = false)
    private int bathrooms;

    @Column(name = "monthly_rent", nullable = false)
    private double monthlyRent;

    @Column(nullable = false, length = 3)
    private String currency = "USD";

    @Column(nullable = false)
    private double deposit = 0.0;

    @Column(nullable = false, length = 255)
    private String address;

    @Column(nullable = false, length = 80)
    private String city;

    @Column(nullable = false, length = 80)
    private String country;

    @Column(nullable = false)
    private double latitude;

    @Column(nullable = false)
    private double longitude;

    @Column(name = "available_from")
    private LocalDate availableFrom;

    @Column(nullable = false, length = 20)
    private String leaseTerm;

    @Column(nullable = false)
    private boolean furnished = false;

    @Column(name = "includes_utilities", nullable = false)
    private boolean includesUtilities = false;

    @Column(name = "image_urls", length = -1)
    private String imageUrls;

    @Column(nullable = false)
    private boolean isActive = true;

    @Column(name = "application_count", nullable = false)
    private int applicationCount = 0;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @LastModifiedDate
    @Column(nullable = false)
    private LocalDateTime updatedAt;
}
