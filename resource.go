package gam

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"io/ioutil"
	"os"
	"strings"
)

// Singleton
type ResourceManager struct {
	shaders  map[string]*Shader
	textures map[string]*Texture2D
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		shaders:  map[string]*Shader{},
		textures: map[string]*Texture2D{},
	}
}

func (r *ResourceManager) LoadShadersPath(path string) ([]*Shader, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var shaders []*Shader
	for i := 0; i < len(files); i+=2 {
		name := strings.Split(files[i].Name(), ".")[0]
		shader, err := r.LoadShader(path + files[i+1].Name(), path + files[i].Name(), name)
		if err != nil {
			return nil, err
		}
		fmt.Println("Loaded shader", name)
		shaders = append(shaders, shader)
	}
	return shaders, nil
}

func (r *ResourceManager) LoadShader(vertexPath, fragmentPath, name string) (*Shader, error) {
	var vertexCode, fragmentCode string

	{
		bytes, err := ioutil.ReadFile(vertexPath)
		if err != nil {
			return nil, err
		}
		vertexCode = string(bytes)

		bytes, err = ioutil.ReadFile(fragmentPath)
		if err != nil {
			return nil, err
		}
		fragmentCode = string(bytes)
	}

	shader := NewShader(vertexCode, fragmentCode)
	if _, ok := r.shaders[name]; ok {
		return nil, fmt.Errorf("shader name already taken")
	}
	r.shaders[name] = shader
	return shader, nil
}

func (r *ResourceManager) Shader(name string) *Shader {
	shader, ok := r.shaders[name]
	if !ok {
		panic("Shader not found")
	}
	return shader
}

func (r *ResourceManager) LoadTexturesPath(path string) ([]*Texture2D, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var textures []*Texture2D
	for i := 0; i < len(files); i++ {
		name := strings.Split(files[i].Name(), ".")[0]
		texture, err := r.LoadTexture(path + files[i].Name(), name)
		if err != nil {
			return nil, err
		}
		fmt.Println("Loaded texture", name)
		textures = append(textures, texture)
	}
	return textures, nil
}

func (r *ResourceManager) LoadTexture(file string, name string) (*Texture2D, error) {
	texture := NewTexture()
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	texture.Generate(f)
	if _, ok := r.textures[name]; ok {
		return nil, fmt.Errorf("texture name already taken")
	}
	r.textures[name] = texture
	return texture, nil
}

func (r *ResourceManager) Texture(name string) *Texture2D {
	t, ok := r.textures[name]
	if !ok {
		panic("Texture '" + name + "' not found")
	}
	return t
}

func (r *ResourceManager) Clear() {
	for _, shader := range r.shaders {
		gl.DeleteProgram(shader.ID)
	}
	for _, texture := range r.textures {
		gl.DeleteTextures(1, &texture.ID)
	}
}
