package com.samuelssantos;

import io.quarkus.test.junit.QuarkusTest;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;
import static org.hamcrest.CoreMatchers.is;

@QuarkusTest
public class ProductResourceTest {

    @Test
    public void testGetProduct() {
        given()
                .when().get("/product")
                .then()
                .statusCode(200);
    }

}