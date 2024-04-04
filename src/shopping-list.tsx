import NewRecipeModal from "./new-recipe-modal";
import DateRangePicker from './DatePicker';

const ShoppingList = () => {
    return <div>
        <div className="mx-8">
            <DateRangePicker />
            <button className="border rounded p-2 bg-green-300 hover:bg-green-400">Generate shopping list</button>
        </div>
        <div className="border border-green-500 rounded mx-8">shopping list</div>
        <div className="mx-8">
            <NewRecipeModal />
        </div> 
    </div>;
  };
  
  export default ShoppingList;