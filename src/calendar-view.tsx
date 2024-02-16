import leftArrow from './assets/left-arrow.svg';
import rightArrow from './assets/right-arrow.svg';
import siteLogo from './assets/siteLogo.png';

const Calendar = () => {

    const getDayWithSuffix = (day) => {
        if (day >= 11 && day <= 13) {
          return `${day}th`;
        }
      
        switch (day % 10) {
          case 1:
            return `${day}st`;
          case 2:
            return `${day}nd`;
          case 3:
            return `${day}rd`;
          default:
            return `${day}th`;
        }
      };

    const currentDate = new Date();

    // Get the current month
    const monthOptions = { month: 'long' };
    const formattedMonth = new Intl.DateTimeFormat('en-US', monthOptions).format(currentDate);

    const currentDay = currentDate.getDay(); // 0 for Sunday, 1 for Monday, ..., 6 for Saturday

    const difference = currentDay - 1 < 0 ? 6 : currentDay - 1;
    const startOfWeek = new Date(currentDate);
    startOfWeek.setDate(currentDate.getDate() - difference);

    const datesOfWeekWithSuffix = [];

    for (let i = 0; i < 7; i++) {
      const date = new Date(startOfWeek);
      date.setDate(startOfWeek.getDate() + i);
      const dayWithSuffix = getDayWithSuffix(date.getDate());
      datesOfWeekWithSuffix.push(date.toLocaleDateString('en-US', { weekday: 'long' }) + ' ' + dayWithSuffix);
    }
    console.log(datesOfWeekWithSuffix);

    return <div className="">
        <div className="flex text-green-500 mb-4">
            <img src={leftArrow} alt="left arrow" className="mr-2" />
            <div className="flex items-center">{formattedMonth}</div>
            <ul className="flex justify-between">{datesOfWeekWithSuffix.map((date, index) => <li key={index} className="border rounded-xl border-green-500 mx-2 px-2">{date}</li>)}</ul>
            <img src={rightArrow} alt="left arrow" />
        </div>
        <div className="flex bg-green-100 h-64 items-center overflow-y-auto mb-4">
            <div className="ml-48 mr-12 w-36">
                <div className="text-2xl text-green-500 mb-4">Breakfast</div>
                <button className="bg-green-500 rounded-full px-4 text-white">+ Add</button>
            </div>
            {/* map the breakfast recipes and for each */}
            <div className="flex-none h-5/6 bg-green-300 rounded-xl mr-8">
                <img className="h-3/4 rounded-t-xl" src={siteLogo} alt="recipe pic" />
                <div className="h-1/4 flex justify-between items-center text-green-300">
                    <div className="text-green-700 ml-4 flex items-center">Name</div>
                    <button className="bg-green-500 text-white mr-4 px-4 rounded-xl">Delete</button>
                </div>
            </div>
        </div>
        <div className="flex bg-green-100 h-64 items-center overflow-y-auto mb-4">
            <div className="ml-48 w-36 mr-12">
                <div className="text-2xl text-green-500 mb-4">Lunch</div>
                <button className="bg-green-500 rounded-full px-4 text-white">+ Add</button>
            </div>
            {/* map the breakfast recipes and for each */}
            <div className="flex-none h-5/6 bg-green-300 rounded-xl mr-8">
                <img className="h-3/4 rounded-t-xl" src={siteLogo} alt="recipe pic" />
                <div className="h-1/4 flex justify-between items-center text-green-300">
                    <div className="text-green-700 ml-4 flex items-center">Name</div>
                    <button className="bg-green-500 text-white mr-4 px-4 rounded-xl">Delete</button>
                </div>
            </div>
        </div>
        <div className="flex bg-green-100 h-64 items-center overflow-y-auto mb-4">
            <div className="ml-48 w-36 mr-12">
                <div className="text-2xl text-green-500 mb-4">Dinner</div>
                <button className="bg-green-500 rounded-full px-4 text-white">+ Add</button>
            </div>
            {/* map the breakfast recipes and for each */}
            <div className="flex-none h-5/6 bg-green-300 rounded-xl mr-8">
                <img className="h-3/4 rounded-t-xl" src={siteLogo} alt="recipe pic" />
                <div className="h-1/4 flex justify-between items-center text-green-300">
                    <div className="text-green-700 ml-4 flex items-center">Name</div>
                    <button className="bg-green-500 text-white mr-4 px-4 rounded-xl">Delete</button>
                </div>
            </div>
            <div className="flex-none h-5/6 bg-green-300 rounded-xl mr-8">
                <img className="h-3/4 rounded-t-xl" src={siteLogo} alt="recipe pic" />
                <div className="h-1/4 flex justify-between items-center text-green-300">
                    <div className="text-green-700 ml-4 flex items-center">Name</div>
                    <button className="bg-green-500 text-white mr-4 px-4 rounded-xl">Delete</button>
                </div>
            </div>
            <div className="flex-none h-5/6 bg-green-300 rounded-xl mr-8">
                <img className="h-3/4 rounded-t-xl" src={siteLogo} alt="recipe pic" />
                <div className="h-1/4 flex justify-between items-center text-green-300">
                    <div className="text-green-700 ml-4 flex items-center">Name</div>
                    <button className="bg-green-500 text-white mr-4 px-4 rounded-xl">Delete</button>
                </div>
            </div>
            <div className="flex-none h-5/6 bg-green-300 rounded-xl mr-8">
                <img className="h-3/4 rounded-t-xl" src={siteLogo} alt="recipe pic" />
                <div className="h-1/4 flex justify-between items-center text-green-300">
                    <div className="text-green-700 ml-4 flex items-center">Name</div>
                    <button className="bg-green-500 text-white mr-4 px-4 rounded-xl">Delete</button>
                </div>
            </div>
            <div className="flex-none h-5/6 bg-green-300 rounded-xl mr-8">
                <img className="h-3/4 rounded-t-xl" src={siteLogo} alt="recipe pic" />
                <div className="h-1/4 flex justify-between items-center text-green-300">
                    <div className="text-green-700 ml-4 flex items-center">Name</div>
                    <button className="bg-green-500 text-white mr-4 px-4 rounded-xl">Delete</button>
                </div>
            </div>
        </div>
       
    </div>;
  };
  
  export default Calendar;