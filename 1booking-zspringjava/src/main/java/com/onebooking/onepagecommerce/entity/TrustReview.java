package com.onebooking.onepagecommerce.entity;

import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;

import java.time.LocalDateTime;

@Entity
@Table(name = "trust_reviews")
@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class TrustReview {

    @Id
    @Column(length = 36, updatable = false)
    private String id;

    @Column(name = "reviewer_id", length = 36, nullable = false)
    private String reviewerId;

    @Column(name = "subject_id", length = 36, nullable = false)
    private String subjectId;

    @Column(name = "subject_type", nullable = false, length = 20)
    private String subjectType;

    @Column(nullable = false)
    private int rating;

    @Column(length = 2000)
    private String comment;

    @Column(nullable = false)
    private boolean isVerified = false;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt;
}
