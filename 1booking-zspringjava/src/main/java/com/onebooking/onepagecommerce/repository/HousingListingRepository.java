package com.onebooking.onepagecommerce.repository;

import com.onebooking.onepagecommerce.entity.HousingListing;
import org.springframework.data.domain.*;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.*;

public interface HousingListingRepository extends JpaRepository<HousingListing, String> {
    Page<HousingListing> findByCityContainingIgnoreCase(String city, Pageable pageable);
    Page<HousingListing> findByCountry(String country, Pageable pageable);
    Page<HousingListing> findByPropertyType(String propertyType, Pageable pageable);
    Page<HousingListing> findByMonthlyRentBetween(double min, double max, Pageable pageable);
    Page<HousingListing> findByLandlordId(String landlordId, Pageable pageable);
    Page<HousingListing> findByPropertyTypeAndMonthlyRentBetween(String type, double min, double max, Pageable pageable);
}
