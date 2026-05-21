package com.onebooking.onepagecommerce.dto;

import com.onebooking.onepagecommerce.entity.User;
import com.onebooking.onepagecommerce.entity.UserDto;
import com.onebooking.onepagecommerce.util.Helper;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;

import java.time.format.DateTimeFormatter;

public class AuthMapper {

    public static UserDto.UserPublicDto toUserPublicDto(User user) {
        return UserDto.UserPublicDto.builder()
                .id(user.getId())
                .email(user.getEmail())
                .firstName(user.getFirstName())
                .lastName(user.getLastName())
                .role(user.getRole())
                .isVerified(user.isVerified())
                .verificationLevel(user.getVerificationLevel())
                .trustScore(user.getTrustScore())
                .avatarUrl(user.getAvatarUrl())
                .createdAt(user.getCreatedAt())
                .updatedAt(user.getUpdatedAt())
                .build();
    }

    public static User toEntity(RegisterRequest req, String passwordHash) {
        return User.builder()
                .id(java.util.UUID.randomUUID().toString())
                .email(req.getEmail())
                .phone(req.getPhone())
                .firstName(req.getFirstName())
                .lastName(req.getLastName())
                .role(req.getRole() != null ? req.getRole() : "DIASPORA")
                .passwordHash(passwordHash)
                .createdAt(java.time.LocalDateTime.now())
                .updatedAt(java.time.LocalDateTime.now())
                .build();
    }
}
