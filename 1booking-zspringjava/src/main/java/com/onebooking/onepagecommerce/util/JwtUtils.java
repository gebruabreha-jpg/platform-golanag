package com.onebooking.onepagecommerce.util;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.io.Decoders;
import io.jsonwebtoken.security.Keys;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Component;

import javax.crypto.SecretKey;
import java.security.Key;
import java.util.*;
import java.util.function.Function;

@Component
public class JwtUtils {

    private final String jwtSecret;
    private final String jwtRefreshSecret;
    private final long accessExpiryMinutes;
    private final long refreshExpiryDays;
    private SecretKey accessKey;
    private SecretKey refreshKey;

    public JwtUtils(
            jakarta.annotation.PostConstruct void init) {
    }
}
