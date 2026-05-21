package com.onebooking.onepagecommerce.config;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.onebooking.onepagecommerce.entity.User;
import com.onebooking.onepagecommerce.repository.UserRepository;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.*;
import lombok.RequiredArgsConstructor;
import org.springframework.http.*;
import org.springframework.security.authentication.*;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;
import java.util.*;

@Component
@RequiredArgsConstructor
public class JwtAuthenticationFilter extends OncePerRequestFilter {

    private final JwtUtils jwtUtils;
    private final UserRepository userRepository;
    private final ObjectMapper objectMapper;

    @Override
    protected void doFilterInternal(HttpServletRequest request,
                                    HttpServletResponse response,
                                    FilterChain filterChain) throws ServletException, IOException {
        try {
            final String authHeader = request.getHeader("Authorization");
            if (authHeader == null || !authHeader.startsWith("Bearer ")) {
                filterChain.doFilter(request, response);
                return;
            }
            String jwt = authHeader.substring(7);
            String username;
            try {
                username = jwtUtils.extractUsername(jwt);
            } catch (Exception e) {
                sendAuthError(response, "Invalid or expired token.");
                return;
            }
            if (username != null && SecurityContextHolder.getContext().getAuthentication() == null) {
                User user = userRepository.findByEmail(username).orElse(null);
                if (user == null || !user.isActive()) {
                    sendAuthError(response, "User not found or disabled.");
                    return;
                }
                if (jwtUtils.isAccessTokenValid(jwt, org.springframework.security.core.userdetails.User
                        .withUsername(user.getEmail())
                        .password(user.getPasswordHash())
                        .authorities(new SimpleGrantedAuthority("ROLE_" + user.getRole()))
                        .build())) {
                    UsernamePasswordAuthenticationToken auth =
                            new UsernamePasswordAuthenticationToken(user, null, List.of(
                                    new SimpleGrantedAuthority("ROLE_" + user.getRole())));
                    auth.setDetails(user);
                    SecurityContextHolder.getContext().setAuthentication(auth);
                }
            }
            filterChain.doFilter(request, response);
        } catch (Exception e) {
            sendAuthError(response, "Authentication failed: " + e.getMessage());
        }
    }

    private void sendAuthError(HttpServletResponse response, String message) throws IOException {
        response.setContentType(MediaType.APPLICATION_JSON_VALUE);
        response.setCharacterEncoding("UTF-8");
        response.setStatus(HttpServletResponse.SC_UNAUTHORIZED);
        Map<String, Object> body = Map.of(
                "success", false,
                "error", Map.of("code", "UNAUTHORIZED", "message", message)
        );
        objectMapper.writeValue(response.getWriter(), body);
    }
}
