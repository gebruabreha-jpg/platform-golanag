package com.onebooking.onepagecommerce.entity;

import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;

import java.time.LocalDateTime;

@Entity
@Table(name = "jobs")
@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class Job {

    @Id
    @Column(length = 36, updatable = false)
    private String id;

    @Column(name = "employer_id", length = 36, nullable = false)
    private String employerId;

    @Column(nullable = false, length = 200)
    private String title;

    @Column(nullable = false, length = 5000)
    private String description;

    @Column(nullable = false, length = 20)
    private String jobType;

    @Column(nullable = false)
    private boolean remote = false;

    @Column(length = 120)
    private String location;

    @Column(length = 80)
    private String country;

    @Column(name = "salary_min")
    private Double salaryMin;

    @Column(name = "salary_max")
    private Double salaryMax;

    @Column(nullable = false, length = 3)
    private String currency = "USD";

    @Column(length = 80)
    private String industry;

    @Column(length = -1)
    private String skills;

    @Column(length = -1)
    private String benefits;

    @Column(name = "application_url")
    private String applicationUrl;

    private LocalDateTime expiresAt;

    @Column(nullable = false)
    private boolean isActive = true;

    @Column(nullable = false)
    private int viewCount = 0;

    @Column(name = "application_count", nullable = false)
    private int applicationCount = 0;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @LastModifiedDate
    @Column(nullable = false)
    private LocalDateTime updatedAt;
}
