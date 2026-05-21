package com.onebooking.onepagecommerce.config;

import io.jsonwebtoken.*;
import io.jsonwebtoken.io.Decoders;
import io.jsonwebtoken.security.Keys;
import jakarta.annotation.PostConstruct;
import lombok.RequiredArgsConstructor;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Component;

import javax.crypto.SecretKey;
import java.security.Key;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.function.Function;

@Component
@RequiredArgsConstructor
public class JwtUtils {

    private final JwtConfig jwtConfig;

    private SecretKey accessKey;
    private SecretKey refreshKey;

    @PostConstruct
    public void init() {
        accessKey  = Keys.hmacShaKeyFor(Decoders.BASE64.decode(base64Encode(jwtConfig.getSecret())));
        refreshKey = Keys.hmacShaKeyFor(Decoders.BASE64.decode(base64Encode(jwtConfig.getRefreshSecret())));
    }

    private static String base64Encode(String s) {
        return java.util.Base64.getEncoder().encodeToString(s.getBytes());
    }

    public String extractUsername(String token) {
        return extractClaim(token, Claims::getSubject);
    }

    public <T> T extractClaim(String token, Function<Claims, T> claimsResolver) {
        final Claims claims = extractAllClaims(token);
        return claimsResolver.apply(claims);
    }

    private Claims extractAllClaims(String token) {
        return Jwts.parser()
                .verifyWith(accessKey)
                .build()
                .parseSignedClaims(token)
                .getPayload();
    }

    public String generateAccessToken(String username, String role) {
        return buildToken(username, role, jwtConfig.getAccessExpiryMinutes() * 60 * 1000, accessKey);
    }

    public String generateRefreshToken(String username) {
        return buildToken(username, null, jwtConfig.getRefreshExpirySeconds() * 1000, refreshKey);
    }

    private String buildToken(String username, String role, long ttlMs, Key key) {
        Map<String, Object> claims = new HashMap<>();
        if (role != null) claims.put("role", role);
        return Jwts.builder()
                .claims(claims)
                .subject(username)
                .issuedAt(new Date())
                .expiration(new Date(System.currentTimeMillis() + ttlMs))
                .signWith(key)
                .compact();
    }

    public boolean isAccessTokenValid(String token, UserDetails userDetails) {
        try {
            final String username = extractUsername(token);
            return (username.equals(userDetails.getUsername())) && !isTokenExpired(token);
        } catch (JwtException | IllegalArgumentException e) {
            return false;
        }
    }

    public boolean isRefreshTokenValid(String token) {
        try {
            Jwts.parser().verifyWith(refreshKey).build().parseSignedClaims(token);
            return true;
        } catch (JwtException | IllegalArgumentException e) {
            return false;
        }
    }

    private boolean isTokenExpired(String token) {
        return extractExpiration(token).before(new Date());
    }

    private Date extractExpiration(String token) {
        return extractClaim(token, Claims::getExpiration);
    }

    public SecretKey getAccessKey()  { return accessKey; }
    public SecretKey getRefreshKey() { return refreshKey; }
    public long getRefreshExpirySeconds() { return jwtConfig.getRefreshExpirySeconds(); }
}
