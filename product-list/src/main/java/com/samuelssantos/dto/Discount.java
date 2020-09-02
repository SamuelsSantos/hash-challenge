package com.samuelssantos.dto;

public class Discount {
    private Float percent;
    private long valueInCents;

    public Float getPercent() {
        return percent;
    }

    public void setPercent(Float percent) {
        this.percent = percent;
    }

    public long getValueInCents() {
        return valueInCents;
    }

    public void setValueInCents(long valueInCents) {
        this.valueInCents = valueInCents;
    }
}
