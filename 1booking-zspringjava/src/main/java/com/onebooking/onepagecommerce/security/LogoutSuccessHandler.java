package com.onebooking.onepagecommerce.security;

import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.security.core.Authentication;
import org.springframework.security.web.authentication.logout.*;
import lombok.RequiredArgsConstructor;
import org.springframework.data.redis.core.RedisTemplate;
import java.util.UUID;

@RequiredArgsConstructor
public class LogoutSuccessHandler extends AbstractAuthenticationTargetUrlRequestHandler
        implements LogoutSuccessHandler {

    private final RedisTemplate<String, String> redisTemplate;

    @Override
    public void onLogoutSuccess(HttpServletRequest request,
                                HttpServletResponse response,
                                Authentication authentication) throws java.io.IOException {
        if (authentication != null) {
            redisTemplate.delete("jwt:blacklist:" + authentication.getCredentials());
        }
        response.setStatus(HttpServletResponse.SC_OK);
        response.getWriter().write("{\"success\":true,\"message\":\"Logged out successfully\"}");
    }
}
