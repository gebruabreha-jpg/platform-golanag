package com.onebooking.onepagecommerce.repository;

import com.onebooking.onepagecommerce.entity.MarketplaceItem;
import org.springframework.data.domain.*;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.*;

public interface MarketplaceItemRepository extends JpaRepository<MarketplaceItem, String> {
    Page<MarketplaceItem> findBySellerId(String sellerId, Pageable pageable);
    Page<MarketplaceItem> findByCategory(String category, Pageable pageable);
    Page<MarketplaceItem> findByCondition(String condition, Pageable pageable);
    Page<MarketplaceItem> findByCategoryAndCondition(String category, String condition, Pageable pageable);
    Page<MarketplaceItem> findByPriceBetween(double min, double max, Pageable pageable);
    Page<MarketplaceItem> findByCountry(String country, Pageable pageable);
    Page<MarketplaceItem> findByCategoryAndPriceBetween(String category, double min, double max, Pageable pageable);
    Page<MarketplaceItem> findByCategoryContainingIgnoreCaseOrTitleContainingIgnoreCase(String category, String title, Pageable pageable);
}
