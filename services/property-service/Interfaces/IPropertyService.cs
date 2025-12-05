using property_service.Models;

namespace property_service.Interfaces;

public interface IPropertyService
{
    public List<Property> GetProperties();
    public Property GetPropertyById(int? propertyId);
    public int InsertProperty(PropertyCreateRequestDTO propertyCreateRequestDTO);
    public int UpdateProperty(PropertyUpdateRequestDTO propertyUpdateRequestDTO);
    public void DeletePropertyById(int? propertyId);

}
