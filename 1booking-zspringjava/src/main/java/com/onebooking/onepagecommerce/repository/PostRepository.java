package com.onebooking.onepagecommerce.repository;

import com.onebooking.onepagecommerce.entity.Post;
import org.springframework.data.domain.*;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.*;

public interface PostRepository extends JpaRepository<Post, String> {
    List<Post> findByCommunityIdOrderByCreatedAtDesc(String communityId);
    Page<Post> findByUserId(String userId, Pageable pageable);
    Page<Post> findByType(String type, Pageable pageable);
}
