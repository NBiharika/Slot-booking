package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginAndRegister(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "LoginAndRegister"})
}

//func ViewSlots(ctx *gin.Context) {
//	ctx.HTML(http.StatusOK, "slot.html", gin.H{"title": "view-slots"})
//}

//login -> give email and password ->create auth token
//register -> add user
//<p>13/07/2022</p>
//<div class="app-check">
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">10:00
//</label>
//</div>
//<input type="radio" class="option-input radio"name="example" />
//<div class="app-border">
//
//<label class="app-label">10:30
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example"/>
//<div class="app-border">
//
//<label class="app-label">11:00
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">11:30
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">12:00
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">12:30
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">13:00
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">13:30
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">14:00
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">14:30
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">15:00
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">15:30
//</label>
//</div>
//</div>
//</div>
//<div>
//<div class="app-check">
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">16:00
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">16:30
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">17:00
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">17:30
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">18:00
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">18:30
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">19:00
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">19:30
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">20:00
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">20:30
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">21:00
//</label>
//</div>
//<input type="radio" class="option-input radio" name="example" />
//<div class="app-border">
//
//<label class="app-label">21:30
//</label>
//</div>
//</div>

//<div class="userLogin">
//<div class="app-time">
//<div>
//<input type="radio" class="option-input radio" name="example"/>
//<div class="app-border">
//<label class="app-label">10:00</label>
//</div>
//{{range .slots}}
//<input type="radio" class="option-input radio" name="example"/>
//<div class="app-border">
//<label class="app-label">{{.StartTime}}</label>
//</div>
//{{end}}
//</div>
//</div>
//</div>
//</div>
