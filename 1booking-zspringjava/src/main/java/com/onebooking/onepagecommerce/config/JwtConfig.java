package com.onebooking.onepagecommerce.config;

import lombok.*;
import org.springframework.boot.context.properties.*;
import org.springframework.context.annotation.*;

@Configuration
@EnableConfigurationProperties
@Data
@ConfigurationProperties(prefix = "jwt")
public class JwtConfig {

    private String secret;
    @Value("${jwt.refresh-secret:}")
    private String refreshSecret;
    private long accessExpiryMinutes = 15;
    private long refreshExpiryDays = 7;

    public long getRefreshExpirySeconds() {
        return refreshExpiryDays * 24 * 60 * 60;
    }
}
