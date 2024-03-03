import NewRecipeModal from "./new-recipe-modal";
import DateRangePicker from './DatePicker';

const ShoppingList = () => {
    return <div>
        <div>calendar from</div>
        <div>calendar to</div>
        <DateRangePicker />
        <div className="border border-green-500 rounded mx-8">shopping list</div>
        <NewRecipeModal />
    </div>;
  };
  
  export default ShoppingList;