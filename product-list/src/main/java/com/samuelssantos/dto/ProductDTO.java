package com.samuelssantos.dto;

import javax.json.bind.annotation.JsonbProperty;
import javax.json.bind.annotation.JsonbPropertyOrder;
import javax.json.bind.config.PropertyOrderStrategy;

@JsonbPropertyOrder(PropertyOrderStrategy.ANY)
public class ProductDTO {

    private String id;

    @JsonbProperty("price_in_cents")
    private long priceInCents;
    private String title;
    private String description;
    @JsonbProperty("Discount")
    private DiscountDTO discountDTO;

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public long getPriceInCents() {
        return priceInCents;
    }

    public void setPriceInCents(long priceInCents) {
        this.priceInCents = priceInCents;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public DiscountDTO getDiscount() {
        return discountDTO;
    }

    public void setDiscount(DiscountDTO discountDTO) {
        this.discountDTO = discountDTO;
    }


}
