package start_up

import (
	"Slot_booking/controller"
	"Slot_booking/manager"
	"Slot_booking/service"
)

var (
	slotRepository manager.SlotRepository
	slotService    service.SlotService
	SlotController controller.SlotController

	userRepository  manager.UserRepository
	userService     service.UserService
	UserController  controller.UserController
	TokenController controller.TokenController

	bookingRepository manager.BookingRepository
	bookingService    service.BookingService
	BookingController controller.BookingController
)

func Initialize() {
	manager.InitializeDB()
	slotRepository = manager.SlotRepo()
	slotService = service.NewSlotService(slotRepository)
	SlotController = controller.NewSlotController(slotService)

	userRepository = manager.UserRepo()
	userService = service.NewUserService(userRepository)
	UserController = controller.NewUserController(userService)
	TokenController = controller.NewTokenController(userService)

	bookingRepository = manager.BookingRepo()
	bookingService = service.NewService(bookingRepository)
	BookingController = controller.New(bookingService, slotService, userService)

}
