package gcp

import (
	"context"
	"log"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

const jsPath = "C:/Users/APostescu/Downloads/gcpnetwork-349117-21f2f3fa3c28.json"
const project = "gcpnetwork-349117"

func TestWorkFlow(t *testing.T) {

	convey.Convey("New() test", t, func() {
		c, err := New(context.Background(), project, "JSPATH", log.Default(), log.Default())
		convey.So(c, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Client.CreateInstance() test", t, func() {
		zone := "europe-west4-c"
		c, _ := New(context.Background(), project, jsPath, log.Default(), log.Default())
		sourceImage := "projects/debian-cloud/global/images/family/debian-9"
		err := c.CreateInstance(context.Background(), zone, "instance4", "n1-standard-2", sourceImage, "default", "default")
		convey.So(err, convey.ShouldBeNil)
	})

	convey.Convey("Client.StartOrStopInstance() test", t, func() {
		// Creating an instance
		zone := "europe-west4-c"
		c, _ := New(context.Background(), project, jsPath, log.Default(), log.Default())
		sourceImage := "projects/debian-cloud/global/images/family/debian-9"
		err := c.CreateInstance(context.Background(), zone, "instance6", "n1-standard-2", sourceImage, "pnetwork", "default")
		convey.So(err, convey.ShouldBeNil)

		// Getting the instance
		instance, err := c.GetInstance(context.Background(), zone, "instance6")

		convey.So(err, convey.ShouldBeNil)
		convey.So(instance.Name, convey.ShouldEqual, "instance6")
		convey.So(instance.Status, convey.ShouldEqual, "RUNNING")

		// Stopping the instance

		err = c.StopInstance(context.Background(), zone, instance.Name)

		convey.So(err, convey.ShouldBeNil)
		convey.So(instance.Status, convey.ShouldEqual, "TERMINATED")

	})

}
