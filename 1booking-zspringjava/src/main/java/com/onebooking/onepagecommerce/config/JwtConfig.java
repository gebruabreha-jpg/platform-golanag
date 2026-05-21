package com.onebooking.onepagecommerce.config;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

@Configuration
@Data
@ConfigurationProperties(prefix = "jwt")
public class JwtConfig {

    private String secret;
    private String refreshSecret = "";
    private long accessExpiryMinutes = 15;
    private long refreshExpiryDays = 7;

    public long getRefreshExpirySeconds() {
        return refreshExpiryDays * 24 * 60 * 60;
    }
}
