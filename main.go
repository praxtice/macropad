package main
import (
	"github.com/rakyll/launchpad"
	"github.com/go-vgo/robotgo"
	"log"
	"time"
)

func main() {
	pad, err := launchpad.Open()
	if err != nil {
		log.Fatalf("Error initalizing launchpad: %v", err)
	}
	defer pad.Close()
	defer pad.Clear()
	pad.Clear()

	ch := pad.Listen()
	for {
		select {
		case hit := <-ch:
			pad.Light(hit.X, hit.Y, 3, 3)
			if hit.X == 0 && hit.Y == 0 {
				robotgo.KeyTap("space","command")
			}
			if hit.X == 1 && hit.Y == 0 {
				robotgo.KeyTap("c","control")
			}
			if hit.X == 1 && hit.Y == 1 {
				// robotgo.KeyTap("tab","command")
				robotgo.TypeString("According to all known laws of aviation, there is no way a bee should be able to fly. Its wings are too small to get its fat little body off the ground. The bee, of course, flies anyway because bees don't care what humans think is impossible.")
			}
			if hit.X == 2 && hit.Y == 2 {
				arr := []string{"We're no strangers to love", "You know the rules and so do I", "A full commitment's what I'm thinking of",
					"You wouldn't get this from any other guy", "I just wanna tell you how I'm feeling", "Gotta make you understand", "",
					"Never gonna give you up,", "Never gonna let you down", "Never gonna run around and desert you", "Never gonna make you cry,",
					"Never gonna say goodbye", "Never gonna tell a lie and hurt you"}

				for i := range arr {
					time.Sleep(1 * time.Second)
					robotgo.TypeString(arr[i])
					robotgo.KeyTap("enter")
				}

			}
			time.Sleep(1000 * time.Nanosecond)
			pad.Light(hit.X, hit.Y, 0, 0)
		}
	}

}
