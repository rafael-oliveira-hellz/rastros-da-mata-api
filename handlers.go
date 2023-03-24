package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
	"rastros-da-mata/crud"
	"strconv"
)

type App struct {
	DB     map[string]*mongo.Collection
	Router *mux.Router
}

// createFruitHandler - cria uma nova fruta
func (app *App) createFruitHandler(w http.ResponseWriter, r *http.Request) {
	var fruit crud.Fruit

	err := json.NewDecoder(r.Body).Decode(&fruit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	if err := fruit.Create(context.Background(), app.DB["fruits"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(fruit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// readFruitHandler - lê uma fruta específica usando o ID fornecido
func (app *App) readFruitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var fruit crud.Fruit

	if err := fruit.Read(context.Background(), app.DB["fruits"], id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(fruit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// updateFruitHandler - atualiza uma fruta usando o ID fornecido e os dados do corpo da requisição
func (app *App) updateFruitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var fruit crud.Fruit

	err = json.NewDecoder(r.Body).Decode(&fruit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	fruit.ID = id

	if err := fruit.Update(context.Background(), app.DB["fruits"], id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(fruit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// deleteFruitHandler - exclui uma fruta usando o ID fornecido
func (app *App) deleteFruitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var fruit crud.Fruit
	fruit.ID = id

	if err := fruit.Delete(app.DB["fruits"], id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllFruitsHandler retorna todas as frutas, limitando e pulando resultados com base em parâmetros de consulta
func (app *App) GetAllFruitsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limitParam := query.Get("limit")
	offsetParam := query.Get("offset")

	if limitParam == "" || offsetParam == "" {
		http.Error(w, "Os parâmetros 'limit' e 'offset' são obrigatórios", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		http.Error(w, "Valor inválido para o parâmetro 'limit'", http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(offsetParam)

	if err != nil {
		http.Error(w, "Valor inválido para o parâmetro 'offset'", http.StatusBadRequest)
		return
	}

	fruits, err := crud.GetAllFruits(app.DB["fruits"], int64(limit), int64(offset))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(fruits)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// createVegetableHandler - cria um novo vegetal
func (app *App) createVegetableHandler(w http.ResponseWriter, r *http.Request) {
	var vegetable crud.Vegetable

	err := json.NewDecoder(r.Body).Decode(&vegetable)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	if err := vegetable.Create(context.Background(), app.DB["vegetables"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(vegetable)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// readVegetableHandler - lê um vegetal específico usando o ID fornecido
func (app *App) readVegetableHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var vegetable crud.Vegetable

	if err := vegetable.Read(context.Background(), app.DB["vegetables"], id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(vegetable)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// updateVegetableHandler - atualiza um vegetal usando o ID fornecido e os dados do corpo da requisição
func (app *App) updateVegetableHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var vegetable crud.Vegetable

	err = json.NewDecoder(r.Body).Decode(&vegetable)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	vegetable.ID = id

	if err := vegetable.Update(context.Background(), app.DB["vegetables"], id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(vegetable)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// deleteVegetableHandler - exclui um vegetal usando o ID fornecido
func (app *App) deleteVegetableHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var vegetable crud.Vegetable
	vegetable.ID = id

	if err := vegetable.Delete(app.DB["vegetables"], id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllVegetablesHandler retorna todos os vegetais, limitando e pulando resultados com base em parâmetros de consulta
func (app *App) GetAllVegetablesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limitParam := query.Get("limit")
	offsetParam := query.Get("offset")

	if limitParam == "" || offsetParam == "" {
		http.Error(w, "Os parâmetros 'limit' e 'offset' são obrigatórios", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		http.Error(w, "Valor inválido para o parâmetro 'limit'", http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(offsetParam)

	if err != nil {
		http.Error(w, "Valor inválido para o parâmetro 'offset'", http.StatusBadRequest)
		return
	}

	vegetables, err := crud.GetAllVegetables(app.DB["vegetables"], int64(limit), int64(offset))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(vegetables)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// createGreenHandler - cria um novo vegetal
func (app *App) createGreenHandler(w http.ResponseWriter, r *http.Request) {
	var green crud.Green

	err := json.NewDecoder(r.Body).Decode(&green)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	if err := green.Create(context.Background(), app.DB["greens"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(green)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// readGreenHandler - lê um vegetal específico usando o ID fornecido
func (app *App) readGreenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var green crud.Green

	if err := green.Read(context.Background(), app.DB["greens"], id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(green)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// updateGreenHandler - atualiza um vegetal usando o ID fornecido e os dados do corpo da requisição
func (app *App) updateGreenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var green crud.Green

	err = json.NewDecoder(r.Body).Decode(&green)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	green.ID = id

	if err := green.Update(context.Background(), app.DB["greens"], id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(green)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// deleteGreenHandler - exclui um vegetal usando o ID fornecido
func (app *App) deleteGreenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var green crud.Green
	green.ID = id

	if err := green.Delete(app.DB["greens"], id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllGreensHandler retorna todos os vegetais, limitando e pulando resultados com base em parâmetros de consulta
func (app *App) GetAllGreensHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limitParam := query.Get("limit")
	offsetParam := query.Get("offset")

	if limitParam == "" || offsetParam == "" {
		http.Error(w, "Os parâmetros 'limit' e 'offset' são obrigatórios", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		http.Error(w, "Valor inválido para o parâmetro 'limit'", http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(offsetParam)

	if err != nil {
		http.Error(w, "Valor inválido para o parâmetro 'offset'", http.StatusBadRequest)
		return
	}

	greens, err := crud.GetAllGreens(app.DB["greens"], int64(limit), int64(offset))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(greens)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
