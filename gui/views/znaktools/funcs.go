package znaktools

import (
	"errors"
	"fmt"
	"path/filepath"
	"zupper/repo"

	"github.com/mechiko/utility"
)

func (p *ZnakToolsPage) ExcelUtilisationCodes(id int64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	rp, err := repo.GetRepository()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	dbZnak, err := rp.LockZnak()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer func() {
		if cerr := rp.UnlockZnak(dbZnak); cerr != nil {
			err = errors.Join(err, cerr)
		}
	}()

	codes, err := dbZnak.UtilisationCodes(id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if name, err := utility.DialogSaveFile(utility.Excel, fmt.Sprintf("%d_отчет_нанесения.xlsx", id), ""); err == nil {
		if filepath.Ext(name) != ".xlsx" {
			name += ".xlsx"
		}
		p.Excel(codes, name)
	}
	return nil
}

func (p *ZnakToolsPage) ExcelOrderCodes(id int64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	rp, err := repo.GetRepository()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	dbZnak, err := rp.LockZnak()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer func() {
		if cerr := rp.UnlockZnak(dbZnak); cerr != nil {
			err = errors.Join(err, cerr)
		}
	}()

	codes, err := dbZnak.OrderCodes(id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if name, err := utility.DialogSaveFile(utility.Excel, fmt.Sprintf("%d_заказ.xlsx", id), ""); err == nil {
		if filepath.Ext(name) != ".xlsx" {
			name += ".xlsx"
		}
		p.Excel(codes, name)
	}
	return nil
}
