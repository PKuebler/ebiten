/*
Copyright 2014 Hajime Hoshi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ebiten

import (
	"image"
)

type Game interface {
	Initialize(g GameContext)
	Update() error
	Draw(gr GraphicsContext) error
}

type GameContext interface {
	IsKeyPressed(key Key) bool
	CursorPosition() (x, y int)
	IsMouseButtonPressed(mouseButton MouseButton) bool
	NewRenderTargetID(width, height int, filter Filter) (RenderTargetID, error)
	NewTextureID(img image.Image, filter Filter) (TextureID, error)
}
