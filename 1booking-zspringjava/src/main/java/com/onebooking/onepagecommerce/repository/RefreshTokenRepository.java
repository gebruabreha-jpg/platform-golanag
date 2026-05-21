package com.onebooking.onepagecommerce.repository;

import com.onebooking.onepagecommerce.entity.RefreshToken;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.*;

public interface RefreshTokenRepository extends JpaRepository<RefreshToken, String> {
    Optional<RefreshToken> findByTokenHashAndUserIdAndRevokedFalse(String tokenHash, String userId);
    List<RefreshToken> findByUserIdAndRevokedFalse(String userId);
    void deleteByUserIdAndRevokedTrue(String userId);
}
