package com.onebooking.onepagecommerce.entity;

import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;

import java.time.LocalDateTime;

@Entity
@Table(name = "posts")
@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class Post {

    @Id
    @Column(length = 36, updatable = false)
    private String id;

    @Column(name = "community_id", length = 36, nullable = false)
    private String communityId;

    @Column(name = "user_id", length = 36, nullable = false)
    private String userId;

    @Column(nullable = false, length = 20)
    private String type;

    @Column(nullable = false, length = 120)
    private String title;

    @Column(nullable = false, length = 4000)
    private String content;

    @Column(name = "media_urls", length = -1)
    private String mediaUrls;

    @Column(nullable = false)
    private boolean isPinned = false;

    @Column(nullable = false)
    private boolean isClosed = false;

    @Column(nullable = false)
    private int replyCount = 0;

    @Column(nullable = false)
    private int viewCount = 0;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt;

    @LastModifiedDate
    @Column(nullable = false)
    private LocalDateTime updatedAt;
}
