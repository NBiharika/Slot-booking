package start_up

import (
	"Slot_booking/cache"
	"Slot_booking/controller"
	"Slot_booking/manager"
	"Slot_booking/service"
)

var (
	slotRepository manager.SlotRepository
	slotService    service.SlotService
	slotCache      cache.SlotCache
	SlotController controller.SlotController

	userRepository  manager.UserRepository
	userService     service.UserService
	userCache       cache.UserCache
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
	slotCache = cache.NewRedisCacheSlot("localhost:8080", 20, cache.OneMonth)
	SlotController = controller.NewSlotController(slotService, slotCache)

	userRepository = manager.UserRepo()
	userService = service.NewUserService(userRepository)
	userCache = cache.NewRedisCache("localhost:8080", 20, cache.OneMonth)
	UserController = controller.NewUserController(userService, userCache)
	TokenController = controller.NewTokenController(userService)

	bookingRepository = manager.BookingRepo()
	bookingService = service.NewService(bookingRepository)
	BookingController = controller.New(bookingService, slotService, userService, userCache)
}
