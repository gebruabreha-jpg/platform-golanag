package com.onebooking.onepagecommerce.repository;

import com.onebooking.onepagecommerce.entity.User;
import org.springframework.data.domain.*;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.*;

public interface UserRepository extends JpaRepository<User, String> {
    Optional<User> findByEmail(String email);
    Optional<User> findByPhone(String phone);
    Optional<User> findByVerificationToken(String token);
    boolean existsByEmail(String email);
    boolean existsByPhone(String phone);
    Page<User> findByRole(String role, Pageable pageable);
}
