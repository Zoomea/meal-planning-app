import NewRecipeModal from "./new-recipe-modal";
import DateRangePicker from './DatePicker';

const ShoppingList = () => {
    return <div>
        <div className="mx-8">
            <DateRangePicker />
        </div>
        <div className="border border-green-500 rounded mx-8">shopping list</div>
        <div className="mx-8">
            <NewRecipeModal />
        </div>
    </div>;
  };
  
  export default ShoppingList;