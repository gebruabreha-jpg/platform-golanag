package com.onebooking.onepagecommerce.entity;

import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;

import java.time.LocalDateTime;

@Entity
@Table(name = "communities")
@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class Community {

    @Id
    @Column(length = 36, updatable = false)
    private String id;

    @Column(nullable = false, length = 100)
    private String name;

    @Column(length = 1000)
    private String description;

    @Column(nullable = false, length = 30)
    private String category;

    @Column(length = 120)
    private String location;

    @Column(length = 80)
    private String country;

    @Column(nullable = false)
    private boolean isPrivate = false;

    @Column(nullable = false)
    private int memberCount = 0;

    @Column(name = "moderator_id", length = 36, nullable = false)
    private String moderatorId;

    @Column(length = 2000)
    private String rules;

    @Column(name = "image_url")
    private String imageUrl;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @LastModifiedDate
    @Column(nullable = false)
    private LocalDateTime updatedAt;
}
