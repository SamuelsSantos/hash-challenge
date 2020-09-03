package com.samuelssantos;

import com.samuelssantos.dto.Discount;
import com.samuelssantos.dto.Product;
import io.quarkus.grpc.runtime.annotations.GrpcService;
import org.jboss.logging.Logger;

import javax.inject.Inject;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.Set;
import com.samuelssantos.pb.DiscountCalculatorServiceGrpc.DiscountCalculatorServiceBlockingStub;
import com.samuelssantos.pb.DiscountRequest;

@Path("/product")
@Produces(MediaType.APPLICATION_JSON)
public class ProductResource {

    private static final Logger logger = Logger.getLogger(ProductResource.class);

    @Inject
    @GrpcService("discount-calculator")
    DiscountCalculatorServiceBlockingStub serviceBlockingStub;

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
    public Response findAll() {
        return Response.ok(products).build();
    }

    @GET
    @Path("/{id}")
    public Response findById(@PathParam("id") String id) {
        Product product = products
                .stream()
                .filter(p -> p.getId().equals(id))
                .findAny()
                .orElseThrow(NotFoundException::new);
        return Response.ok(product).build();
    }

    @GET
    @Path("/rpc/{id}")
    public void rpc(@PathParam("id") String id) {
        String msg = serviceBlockingStub.process(
                DiscountRequest.newBuilder().setProductId("9090").setUserId(id).build()).getMsg();
        logger.log(Logger.Level.INFO, msg);
    }
}