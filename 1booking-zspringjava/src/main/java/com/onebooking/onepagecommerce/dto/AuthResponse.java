package com.onebooking.onepagecommerce.dto;

import jakarta.validation.constraints.NotBlank;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@AllArgsConstructor @NoArgsConstructor @Builder @Data
public class AuthResponse {
    @NotBlank
    private String accessToken;

    @NotBlank
    private String refreshToken;

    @NotBlank
    private String username;

    @NotBlank
    private String role;

    @NotBlank
    private long expiresAt;

    private UserPublicDto user;
}
