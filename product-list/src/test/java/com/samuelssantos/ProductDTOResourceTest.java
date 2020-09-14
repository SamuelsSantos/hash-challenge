package com.samuelssantos;

import io.quarkus.test.junit.QuarkusTest;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;

@QuarkusTest
public class ProductDTOResourceTest {


    @Test
    public void testDocApi() {
        given()
                .when().get("/swagger-ui")
                .then()
                .statusCode(200);
    }

    @Test
    public void testMicroProfileApi() {
        given()
                .when().get("/health/live")
                .then()
                .statusCode(200);
    }


}