using property_service.Interfaces;
using property_service.Models;

namespace property_service.Services;

public class PropertyService:IPropertyService
{
    public PropertyService()
    {

    }

    public List<Property> GetProperties()
    {
        return [new Property { Id = 1, Name = "Sample Property" }];
    }

    public Property GetPropertyById(int? propertyId)
    {
        return new Property { Id = 1, Name = "Sample Property" };
    }
    public int InsertProperty(PropertyCreateRequestDTO propertyCreateRequestDTO)
    {
        return 1;
    }

    public int UpdateProperty(PropertyUpdateRequestDTO propertyUpdateRequestDTO)
    {
        return 1;
    }
    public void DeletePropertyById(int? propertyId)
    {

    }
}
