using Microsoft.AspNetCore.Diagnostics.HealthChecks;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Diagnostics.HealthChecks;
using property_service.Database;
using property_service.Interfaces;
using property_service.Options;
using property_service.Services;
using Serilog;
using Serilog.Formatting.Compact;
using Serilog.Formatting.Json;
using Serilog.Sinks.Network;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers()
    .AddJsonOptions(options =>
    {
        options.JsonSerializerOptions.ReferenceHandler =
            System.Text.Json.Serialization.ReferenceHandler.IgnoreCycles;
    });
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

//Options
builder.Services
    .AddOptions<SupabaseStorageOptions>()
    .Bind(builder.Configuration.GetSection(SupabaseStorageOptions.SectionName))
    .ValidateDataAnnotations()
    .Validate(opt => !string.IsNullOrEmpty(opt.Url), "Url should not be empty")
    .Validate(opt => !string.IsNullOrEmpty(opt.ServiceRoleKey), "Url should not be empty")
    .ValidateOnStart();

//Dependency Injection
builder.Services.AddScoped<IPropertyService, PropertyService>();
builder.Services.AddSingleton<ISupabaseStorageService, SupabaseStorageService>();

//Database
builder.Services.AddDbContext<PropertyDbContext>(options =>
    options.UseNpgsql(builder.Configuration.GetConnectionString("Supabase")));

//Logging
builder.Host.UseSerilog((context, config) =>
{
    config
        .MinimumLevel.Information()
        .Enrich.FromLogContext()
        .Enrich.WithProperty("Service", "property-service")
        .WriteTo.Console()
        .WriteTo.Http(
            requestUri: "http://localhost:5044", // lokalni Logstash
            queueLimitBytes: null,
            textFormatter: new RenderedCompactJsonFormatter()
        );
});

//Health checks
builder.Services.AddHealthChecks()
    .AddCheck("self", () => HealthCheckResult.Healthy())
    .AddNpgSql(
        builder.Configuration.GetConnectionString("Supabase") ?? "",
        name: "postgres",
        failureStatus: HealthStatus.Unhealthy
    );

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

// Health endpoints
app.MapHealthChecks("/health", new HealthCheckOptions
{
    Predicate = check => check.Name == "self"
});

app.MapHealthChecks("/health/ready", new HealthCheckOptions
{
    Predicate = _ => true
});

app.MapGet("/ok", () =>
{
    Log.Information("OK endpoint called");
    return "OK";
});

app.MapGet("/error", () =>
{
    Log.Error("Something went wrong");
    return Results.Problem("Error");
});

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
