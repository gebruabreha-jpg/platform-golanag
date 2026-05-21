package com.onebooking.onepagecommerce.repository;

import com.onebooking.onepagecommerce.entity.Community;
import org.springframework.data.domain.*;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.*;

public interface CommunityRepository extends JpaRepository<Community, String> {
    List<Community> findByModeratorId(String moderatorId);
    Page<Community> findByCategory(String category, Pageable pageable);
    Page<Community> findByIsPrivate(boolean isPrivate, Pageable pageable);
    Page<Community> findByNameContainingIgnoreCase(String keyword, Pageable pageable);
}
