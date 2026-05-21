package com.onebooking.onepagecommerce.dto;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@AllArgsConstructor @NoArgsConstructor @Builder @Data
public class RegisterRequest {
    @NotBlank @Email
    private String email;

    @NotBlank
    @Size(min = 8, max = 80)
    private String password;

    @NotBlank
    @Size(min = 1, max = 80)
    private String firstName;

    @NotBlank
    @Size(min = 1, max = 80)
    private String lastName;

    @Size(min = 7, max = 30)
    private String phone;

    private String country;
    private String city;
    private String role;
}
