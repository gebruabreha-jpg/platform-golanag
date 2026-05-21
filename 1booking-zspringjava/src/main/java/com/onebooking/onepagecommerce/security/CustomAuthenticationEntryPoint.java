package com.onebooking.onepagecommerce.security;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.onebooking.onepagecommerce.util.Helper;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.web.AuthenticationEntryPoint;
import java.io.IOException;

public final class CustomAuthenticationEntryPoint implements AuthenticationEntryPoint {

    private static final ObjectMapper OM = new ObjectMapper();

    @Override
    public void commence(jakarta.servlet.http.HttpServletRequest request,
                         jakarta.servlet.http.HttpServletResponse response,
                         AuthenticationException authException) throws IOException {
        response.setContentType(MediaType.APPLICATION_JSON_VALUE);
        response.setCharacterEncoding("UTF-8");
        response.setStatus(HttpStatus.UNAUTHORIZED.value());
        OM.writeValue(response.getWriter(), Helper.errorResponse("UNAUTHORIZED", authException.getMessage(), null));
    }
}
