using Microsoft.AspNetCore.Mvc;
using property_service.Interfaces;
using property_service.Models;

[ApiController]
[Route("property")]
[Produces("application/json")]
public class PropertyController : ControllerBase
{
    private readonly IPropertyService _propertyService;

    public PropertyController(IPropertyService propertyService)
    {
        _propertyService = propertyService;
    }

    // GET /Property
    [HttpGet]
    [EndpointSummary("Get all properties")]
    [ProducesResponseType(StatusCodes.Status200OK, Type = typeof(List<Property>))]
    public async Task<ActionResult<List<Property>>> GetProperties()
    {
        var properties = _propertyService.GetProperties();
        return Ok(properties);
    }

    // GET /Property/{id}
    [HttpGet("{id:int}")]
    [EndpointSummary("Retrieves the property matching the specified ID.")]
    [ProducesResponseType(StatusCodes.Status200OK, Type = typeof(Property))]
    [ProducesResponseType(StatusCodes.Status404NotFound)]
    public async Task<ActionResult<Property>> GetById(int id)
    {
        var property = _propertyService.GetPropertyById(id);

        if (property == null)
            return NotFound();

        return Ok(property);
    }

    // POST /Property
    [HttpPost]
    [EndpointSummary("Inserts a new property and returns its ID.")]
    [ProducesResponseType(StatusCodes.Status200OK, Type = typeof(int))]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [Consumes("application/json")]
    public async Task<ActionResult<int>> Insert([FromBody] PropertyCreateRequestDTO createRequestDTO)
    {
        var created = _propertyService.InsertProperty(createRequestDTO);
        return Ok(created);
    }

    // PUT /Property/{id}
    [HttpPut("{id:int}")]
    [EndpointSummary("Updates the property matching the specified ID.")]
    [ProducesResponseType(StatusCodes.Status200OK, Type = typeof(int))]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [Consumes("application/json")]
    public async Task<ActionResult<int>> Update([FromBody] PropertyUpdateRequestDTO updateRequestDTO)
    {
        var id = _propertyService.UpdateProperty(updateRequestDTO);
        return Ok(id);
    }

    // DELETE /Property/{id}
    [HttpDelete("{id:int}")]
    [EndpointSummary("Deletes the property matching the specified ID.")]
    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status404NotFound)]
    public async Task<ActionResult> Delete(int id)
    {
        _propertyService.DeletePropertyById(id);
        return Ok();
    }
}
