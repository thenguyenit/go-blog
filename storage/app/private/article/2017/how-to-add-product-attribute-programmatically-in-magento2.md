#### Step 1: Create a file Setup/InstallData to create product attribute


```
<?php

namespace [VendorName]\[ModuleName]\Setup;

use Magento\Eav\Setup\EavSetup;
use Magento\Eav\Setup\EavSetupFactory;
use Magento\Framework\Setup\InstallDataInterface;
use Magento\Framework\Setup\ModuleContextInterface;
use Magento\Framework\Setup\ModuleDataSetupInterface;

class InstallData implements InstallDataInterface
{

    private $eavSetupFactory;

    public function __construct(EavSetupFactory $eavSetupFactory)
    {
        $this->eavSetupFactory = $eavSetupFactory;
    }

    /**
     * {@inheritdoc}
     */
    public function install(ModuleDataSetupInterface $setup, ModuleContextInterface $context)
    {
        /** @var EavSetup $eavSetup */
        $eavSetup = $this->eavSetupFactory->create(['setup' => $setup]);
        $eavSetup->addAttribute(
            \Magento\Catalog\Model\Product::ENTITY,
            'attribute_code',
            [
                'input' => 'boolean', // Input type
                'type' => 'int',
                'backend' => '',
                'frontend' => '',
                'label' => 'Attribute label',
                'class' => '',
                'source' => '',
                'global' => \Magento\Eav\Model\Entity\Attribute\ScopedAttributeInterface::SCOPE_GLOBAL,
                'visible' => true, //Visible on the backend
                'visible_on_front' => false, //Visible on the front-end
                'required' => false,
                'user_defined' => false,
                'default' => '',
                'searchable' => false,
                'filterable' => false,
                'comparable' => false,
                'used_in_product_listing' => true, //Depends on the theme. To include the attribute in product summaries that appear in catalog listings
                'unique' => false,
                'apply_to' => '' // The product types will be applied
            ]
        );
    }
}
```

#### Step 2: Get custom attribute value

##### Get custom attribute value in the catalog + pdp page


```
<?php

namespace [VendorName]\[ModuleName]\Helper;

use Magento\Store\Model\ScopeInterface;

class Image extends \Magento\Catalog\Helper\Image
{

    /**
     * Current Product
     *
     * @var \Magento\Catalog\Model\Product
     */
    protected $_product;

    public function getYourAttributeValue()
    {
        return $this->_product->getData(attribute_code);
    }
    ...
```

##### Get custom attribute value in cart + checkout page

In Magento 2, the file [VendorName]\[ModuleName]\etc\\**catalog_attributes.xml** is using to define list of attributes will be loaded automatically.


The group quote_item represents the attributes that are going to be copied from the product to the quote item.

Create [VendorName]\[ModuleName]\etc\\**catalog_attributes.xml**


```xml
<?xml version="1.0"?>

<config xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="urn:magento:module:Magento_Catalog:etc/catalog_attributes.xsd">
    <group name="quote_item">
        <attribute name="attribute_code"/>
    </group>
</config>
```

