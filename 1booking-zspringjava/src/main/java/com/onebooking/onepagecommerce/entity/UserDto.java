package com.onebooking.onepagecommerce.entity;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;
import lombok.*;

import java.time.LocalDateTime;

@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class UserDto {

    @NotBlank
    @Email
    private String email;

    @NotBlank
    private String phone;

    @NotBlank
    @Size(min = 1, max = 80)
    private String firstName;

    @NotBlank
    @Size(min = 1, max = 80)
    private String lastName;

    private String bio;
    private String avatarUrl;
    private String location;
    private String country;
    private String city;
    private String role = "DIASPORA";

    @Builder.Default
    private boolean emailVerified = false;

    private LocalDateTime emailVerifiedAt;
    private String verificationToken;
    private String verificationTokenExpiresAt;

    private boolean phoneVerified;

    private String otp;

    private LocalDateTime otpExpiresAt;

    private LocalDateTime lastLoginAt;

    @Builder.Default
    private boolean active = true;

    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;
}
