package com.onebooking.onepagecommerce.entity;

import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import java.time.*;
import java.util.*;

@Entity
@Table(name = "users")
@EntityListeners(AuditingEntityListener.class)
@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class User {

    @Id
    @Column(length = 36, nullable = false, updatable = false)
    private String id;

    @Column(nullable = false, unique = true, length = 255)
    private String email;

    @Column(nullable = false, unique = true, length = 32)
    private String phone;

    @Column(nullable = false, length = 80)
    private String firstName;

    @Column(nullable = false, length = 80)
    private String lastName;

    @Column(name = "date_of_birth")
    private LocalDate dateOfBirth;

    @Column(length = 500)
    private String bio;

    @Column(name = "avatar_url")
    private String avatarUrl;

    @Column(length = 120)
    private String location;

    @Column(length = 80)
    private String country;

    @Column(length = 80)
    private String city;

    @Column(nullable = false, length = 20)
    private String role = "DIASPORA";

    @Column(nullable = false)
    private boolean isVerified = false;

    @Column(name = "verification_level", nullable = false)
    private int verificationLevel = 0;

    @Column(nullable = false)
    private double trustScore = 0.0;

    @Column(name = "total_transactions", nullable = false)
    private int totalTransactions = 0;

    @Column(name = "password_hash", nullable = false, length = 255)
    private String passwordHash;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @LastModifiedDate
    @Column(nullable = false)
    private LocalDateTime updatedAt;

    private LocalDateTime deletedAt;

    @Column(name = "is_deleted", nullable = false)
    private boolean isDeleted = false;

    @OneToMany(mappedBy = "user", cascade = CascadeType.ALL, orphanRemoval = true)
    private List<RefreshToken> refreshTokens;

    public boolean isActive() { return !isDeleted && deletedAt == null; }
}
