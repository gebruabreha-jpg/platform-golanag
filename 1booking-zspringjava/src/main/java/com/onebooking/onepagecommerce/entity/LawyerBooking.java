package com.onebooking.onepagecommerce.entity;

import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;

import java.time.LocalDateTime;

@Entity
@Table(name = "lawyer_bookings")
@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class LawyerBooking {

    @Id
    @Column(length = 36, updatable = false)
    private String id;

    @Column(name = "lawyer_id", length = 36, nullable = false)
    private String lawyerId;

    @Column(name = "user_id", length = 36, nullable = false)
    private String userId;

    @Column(nullable = false)
    private LocalDateTime scheduledAt;

    @Column(nullable = false)
    private int duration;

    @Column(nullable = false, length = 20)
    private String type;

    @Column(nullable = false, length = 20)
    private String status;

    @Column(length = 4000)
    private String notes;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @LastModifiedDate
    @Column(nullable = false)
    private LocalDateTime updatedAt;
}
