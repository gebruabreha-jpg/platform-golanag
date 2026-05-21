package com.onebooking.onepagecommerce.config;

import io.swagger.v3.oas.models.*;
import io.swagger.v3.oas.models.info.*;
import io.swagger.v3.oas.models.security.*;
import io.swagger.v3.oas.models.servers.*;
import org.springdoc.core.customizers.*;
import org.springframework.context.annotation.*;
import org.springframework.security.config.annotation.method.configuration.*;

@Configuration
@EnableMethodSecurity
public class OpenApiConfig {

    @Bean
    public OpenAPI onepagecommerceAPI() {
        return new OpenAPI()
                .info(new Info()
                        .title("1Booking Platform API")
                        .version("1.0.0")
                        .description("""
                                Booking platform REST API – flights, shipping, housing, marketplace, services.
                                Supports JWT access-token + refresh-token authentication with Redis-backed revocation.
                                """))
                .servers(List.of(
                        new Server().url("http://localhost:8080").description("Development"),
                        new Server().url("http://localhost:8080/api/v1").description("API v1")
                ))
                .components(new Components()
                        .addSecuritySchemes("bearerAuth",
                                new SecurityScheme()
                                        .type(SecurityScheme.Type.HTTP)
                                        .scheme("bearer")
                                        .bearerFormat("JWT")
                                        .description("JWT access token obtained from /auth/login")
                        )
                );
    }
}
