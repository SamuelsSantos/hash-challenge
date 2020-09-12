package com.samuelssantos.dto;


import javax.json.bind.annotation.JsonbProperty;
import javax.json.bind.annotation.JsonbPropertyOrder;
import javax.json.bind.config.PropertyOrderStrategy;

@JsonbPropertyOrder(PropertyOrderStrategy.ANY)
public class Discount {

    @JsonbProperty("pct")
    private double percent;

    @JsonbProperty("value_in_cents")
    private long valueInCents;

    public double getPercent() {
        return percent;
    }

    public void setPercent(double percent) {
        this.percent = percent;
    }

    public long getValueInCents() {
        return valueInCents;
    }

    public void setValueInCents(long valueInCents) {
        this.valueInCents = valueInCents;
    }

}
