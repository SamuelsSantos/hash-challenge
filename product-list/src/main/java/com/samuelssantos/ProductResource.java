package com.samuelssantos;

import com.google.protobuf.Empty;
import com.samuelssantos.dto.DiscountDTO;
import com.samuelssantos.dto.ProductDTO;
import com.samuelssantos.pb.DiscountCalculatorServiceGrpc;
import com.samuelssantos.pb.DiscountRequest;
import io.quarkus.cache.CacheResult;
import io.quarkus.grpc.runtime.annotations.GrpcService;
import org.jboss.logging.Logger;
import protorepo.ProductServiceGrpc;
import protorepo.Products;

import javax.inject.Inject;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

@Path("/product")
@Produces(MediaType.APPLICATION_JSON)
public class ProductResource {

    private static final Logger logger = Logger.getLogger(ProductResource.class);
    public static final String EMPTY_USER = "";

    @Inject
    @GrpcService("discountcalculator")
    DiscountCalculatorServiceGrpc.DiscountCalculatorServiceBlockingStub calculatorService;

    @Inject
    @GrpcService("products")
    ProductServiceGrpc.ProductServiceBlockingStub productService;

    public ProductDTO mapToDTO(Products.Product result) {
        ProductDTO productDTO = new ProductDTO();
        DiscountDTO discountDTO = new DiscountDTO();
        productDTO.setDescription(result.getDescription());
        productDTO.setTitle(result.getTitle());
        productDTO.setId(result.getId());
        productDTO.setPriceInCents(result.getPriceInCents());
        discountDTO.setPercent(result.getDiscount().getPct());
        discountDTO.setValueInCents(result.getDiscount().getValueInCents());
        productDTO.setDiscount(discountDTO);
        return productDTO;
    }

    public ProductDTO getDiscount(String userId, Products.Product product) {
        logger.log(Logger.Level.INFO, String.format("User: %s ProductDTO: %s", userId, product.getId()));
        try {
            return mapToDTO(calculatorService.process(buildDiscountRequest(userId, product.getId())).getResult());
        } catch (Exception e) {
            logger.log(Logger.Level.ERROR, e.getMessage());
            return mapToDTO(product);
        }
    }

    private DiscountRequest buildDiscountRequest(String userId, String productId) {
        return DiscountRequest.newBuilder()
                .setProductId(productId)
                .setUserId(userId)
                .build();
    }

    @CacheResult(cacheName = "product-cache")
    private Iterator<Products.Product> getProdutos() {
        Empty request = Empty.newBuilder().build();
        return productService.list(request);
    }

    @GET
    public Response list(@DefaultValue(EMPTY_USER) @QueryParam("X-USER-ID") String userId) {
        List<ProductDTO> result = new ArrayList<>();
        try {
            Iterator<Products.Product> products = getProdutos();
            if (products != null && products.hasNext()) {
                products.forEachRemaining(item -> {
                        result.add(getDiscount(userId, item));
                });
            }

            return Response.ok(result).build();
        } catch (Exception e) {
            logger.log(Logger.Level.ERROR, String.format("RPC failed: %s", e.getMessage()));
            return Response.status(Response.Status.INTERNAL_SERVER_ERROR).build();
        }
    }
}