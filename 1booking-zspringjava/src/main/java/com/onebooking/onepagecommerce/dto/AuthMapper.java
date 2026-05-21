package com.onebooking.onepagecommerce.dto;

import com.onebooking.onepagecommerce.entity.User;
import com.onebooking.onepagecommerce.util.Helper;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;

import java.time.format.DateTimeFormatter;

public class AuthMapper {

    public static UserPublicDto.UserPublicDtoBuilder toUserPublicBuilder(User user) {
        return UserPublicDto.builder()
                .id(user.getId())
                .email(user.getEmail())
                .firstName(user.getFirstName())
                .lastName(user.getLastName())
                .role(user.getRole())
                .isVerified(user.isVerified())
                .verificationLevel(user.getVerificationLevel())
                .trustScore(user.getTrustScore())
                .avatarUrl(user.getAvatarUrl())
                .createdAt(user.getCreatedAt());
    }

    public static User toEntity(RegisterRequest req, String passwordHash) {
        return User.builder()
                .id(Helper.generateUuid())
                .email(req.getEmail())
                .phone(req.getPhone())
                .firstName(req.getFirstName())
                .lastName(req.getLastName())
                .role(req.getRole() != null ? req.getRole() : "DIASPORA")
                .passwordHash(passwordHash)
                .build();
    }
}
