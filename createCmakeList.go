package createta

import (
	"fmt"
	"log"
	"os"
)

func (ctx *Ctactx) WriteCmakeList() {
	cmakelist := `project (%s C)

set (SRC host/main.c)

add_executable (${PROJECT_NAME} ${SRC})

target_include_directories(${PROJECT_NAME}
			   PRIVATE ta/include
			   PRIVATE include)

target_link_libraries (${PROJECT_NAME} PRIVATE teec)

install (TARGETS ${PROJECT_NAME} DESTINATION ${CMAKE_INSTALL_BINDIR})
`
	err := os.WriteFile("CMakeList.txt", []byte(fmt.Sprintf(cmakelist, ctx.lowerName)), MODE)
	if err != nil {
		log.Print("Error creating the CMakeList.txt file: ", err.Error())
	}
}
