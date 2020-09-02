package com.samuelssantos;

import com.samuelssantos.dto.Discount;
import com.samuelssantos.dto.Product;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.Set;

@Path("/product")
@Produces(MediaType.APPLICATION_JSON)
public class ProductResource {

    private Set<Product> products = Collections.newSetFromMap(Collections.synchronizedMap(new LinkedHashMap<>()));

    public ProductResource() {
        for (Integer i = 1; i < 6; i++) {
            Product product = new Product();
            product.setId(i.toString());
            product.setTitle("Product " + i);
            product.setDescription(" Description product " + i);
            product.setDiscount(new Discount());
            product.setPriceInCents(i * 100);
            products.add(product);
        }
    }

    @GET
    public Response list() {
        return Response.ok(products).build();
    }
}