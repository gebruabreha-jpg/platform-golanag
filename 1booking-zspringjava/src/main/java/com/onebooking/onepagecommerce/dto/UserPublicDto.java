package com.onebooking.onepagecommerce.dto;

import io.swagger.v3.oas.annotations.media.Schema;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;

@AllArgsConstructor @NoArgsConstructor @Builder @Data
@Schema(description = "Public user profile (no password hash)")
public class UserPublicDto {
    @Schema(description = "UUID v4 user identifier")
    private String id;

    @Schema(description = "User email address")
    private String email;

    @Schema(description = "First name")
    private String firstName;

    @Schema(description = "Last name")
    private String lastName;

    @Schema(description = "User role: DIASPORA, LOCAL, MERCHANT or ADMIN")
    private String role;

    @Schema(description = "Whether the user's email/phone is verified")
    private boolean isVerified;

    @Schema(description = "Verification level (0-none, 1-basic, 2-full)")
    private int verificationLevel;

    @Schema(description = "Community trust score")
    private double trustScore;

    @Schema(description = "Avatar image URL")
    private String avatarUrl;

    @Schema(description = "ISO-8601 timestamp of account creation")
    private LocalDateTime createdAt;
}
