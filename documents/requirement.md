# Requirements

Company A is a business that need to manage the reliability of energy and gas supply in Singapore. For various projects related to build and maintain the utilities stations, cable drums is one of the most essential parts need to be used. The current processes to procure and consume cable drum are manual with lots of emails and offline coordinations. The involving parties are:

- Planner Team : staff from Company A, to do the capacity planning, procure the cable drum via Contract with Supply Vendors and manage the consumption via coordinating with Sypply Vendors and Project Contractors
- Supply Vendor Team : vendors who have the contract awarded by Company A to supply cable drums, to manufacture and stock cable drums at their own warehouse according to the contracting amount
- Project Contractor Team : vendors who have the contract awarded by Company A to build & maintain the utilities stations, to collect the cable drums from Supply Vendor warehouse and carry out the neccessary project works

There is an initiative to build the web portal which will be able to help various parties to handle the processes end-to-end digitally. The following Use Cases are critical to satisfy:

# Use cases
1. As the System Admin,
   I want to manage the account for users with various roles (Planner Team, Supply Vendor Team and Project Contractor Team)
   So that the users can log in and carry out necessary actions according to their role
   (Assumption: System Admin account is pre-created in DB and cannot be modified, there is no need to do verification during account creation for other users)


2. As the Planner Team user,
   I want to view all the contracts of cable drums supply in the system from all Supply Vendors, each contract will include minimally the start date, end date and contracting amount of cable drums to be supplied by specific Supply Vendor
   So that I can plan and manage the usage of cable drums
   (Assumption: procurement processes involving calling the tenders, selecting vendors/contractors, awarding contracts are not in scope of the system, we only managed the confirmed Contracts, Supply Vendors and Project Contractors. Contracts can be pre-created in DB)


3. As the Planner Team user, I want to plan and manage the usage of cable drums by:
   - creating the Request For Withdraw for cable drums from any available contract with Supply Vendor of choice as long as the quantity not exceeding remaining contracting amount
   - and assigning the newly created Request For Withdraw to the Project Contractor of choice so that the appointed Project Contractor can contact the Supply Vendor to collect the cable drums accordingly
   And then I can see the Request For Withdraw in the system with status "New" together with any existing Request For Withdraw with up-to-date status
   So that I can manage the operations with neccessary actions and solving the issues if needed


4. As the Supply Vendor Team user,
   I want to see the Request For Withdraw from the contract of my team
   So that I can prepare the requesting cable drums in warehouse for collection
   And then I can update the status of Request For Withdraw in the system become "Ready to Collect" when the preparation is finished
   So that the Project Contractor Team can arrange the collection


5. As the Project Contractor Team user,
   I want to receive notification via email when there is new Request For Withdraw assign to my team
   So that I can arrange the collection for cable drum to do neccesary project works
   And then I can log in to the System and update the status of Request For Withdraw become "Collected"
   So that the system can update the remaining cable drums amount of relevant contract for future planning and managing by Planner Team